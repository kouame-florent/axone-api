package server

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/kouame-florent/axone-api/api/grpc/gen"
	"github.com/kouame-florent/axone-api/internal/axone"
	"github.com/kouame-florent/axone-api/internal/config"
	"github.com/kouame-florent/axone-api/internal/repo"
	"github.com/kouame-florent/axone-api/internal/store"
	"github.com/kouame-florent/axone-api/internal/svc"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

const maxAttchmentSize = 20 << 20 //1 Mi

type AxoneServer struct {
	gen.UnimplementedAxoneServer
	subscribers sync.Map
	DB          *gorm.DB
}

type sub struct {
	stream   gen.Axone_SubscribeServer // stream is the server side of the RPC stream
	finished chan<- bool               // finished is used to signal closure of a client subscribing goroutine
}

func NewAxoneServer(db *gorm.DB) *AxoneServer {
	server := &AxoneServer{
		DB: db,
	}

	return server
}

//SendNewTicket send a new ticket to axone
func (s *AxoneServer) SendNewTicket(ctx context.Context, req *gen.NewTicketRequest) (*gen.NewTicketResponse, error) {
	log.Printf("TICKETID: %s", req.TicketID)
	ticketID, err := uuid.Parse(req.TicketID)
	if err != nil {
		return &gen.NewTicketResponse{}, err
	}
	requesterID, err := uuid.Parse(req.RequesterID)
	if err != nil {
		return &gen.NewTicketResponse{}, err
	}

	//db := store.NewDB()
	rep := repo.NewTicketRepo(s.DB)
	ticketSvc := svc.NewTicketSvc(rep)

	id, err := ticketSvc.SendNewTicket(ticketID, req.Subject, req.Request, axone.TicketType(req.Type), requesterID)
	if err != nil {
		return &gen.NewTicketResponse{}, err
	}

	resp := &gen.NewTicketResponse{
		ID: id.String(),
	}

	//notify ax client to refresh tickets list
	s.SendNotification("TICKETS_AVAILABLE")

	return resp, nil
}

//SendAttachment attachment(s) with ticket ID in metadata
func (s *AxoneServer) SendAttachment(stream gen.Axone_SendAttachmentServer) error {
	meta, err := stream.Recv()
	if err != nil {
		zap.L().Error("failed to receive attachment meta", zap.Any("error", err))
		return err
	}

	id := meta.GetInfo().TicketID
	mime := meta.GetInfo().MimeType

	log.Printf("For ticket: %s  with type %s", id, mime)

	atachmentData := bytes.Buffer{}
	attachmentSize := 0

	for {
		log.Print("waiting to receive more data")

		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			zap.L().Error("failed to receive attachment data", zap.Any("error", err))
			return status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err)
		}

		chunk := req.GetChunkData()
		size := len(chunk)

		log.Printf("received a chunk with size: %d", size)

		attachmentSize += size
		if attachmentSize > maxAttchmentSize {
			zap.L().Error("image is too large", zap.Any("size", "15 Mi"))

			return status.Errorf(codes.InvalidArgument, "image is too large: %d > %d", attachmentSize, maxAttchmentSize)
		}
		_, err = atachmentData.Write(chunk)
		if err != nil {
			return status.Errorf(codes.Internal, "cannot write chunk data: %v", err)
		}

	}

	//db := store.NewDB()
	rep := repo.NewAttachmentRepo(s.DB)
	attSvc := svc.NewAttachmentSvc(rep)

	storageName := uuid.New()

	attSvc.RegisterMetas(meta, string(storageName.String()))
	attPath, err := config.AttachmentPath()
	if err != nil {
		zap.L().Error("Cannot get attachment path", zap.Any("error", err))
		return status.Errorf(codes.Internal, "%s", err)
	}

	st := store.NewAttachmentStore(attPath, storageName)
	err = st.StoreFile(atachmentData)
	if err != nil {
		zap.L().Error("Cannot save attachment", zap.Any("error", err))
		return status.Errorf(codes.Internal, "%s", err)
	}

	res := &gen.SendAttachmentResponse{
		TicketID: meta.GetInfo().TicketID,
	}

	err = stream.SendAndClose(res)
	if err != nil {
		zap.L().Error("cannot send response", zap.Any("error", err))
		return status.Errorf(codes.Internal, "%s", err)

	}

	return nil
}

func (s *AxoneServer) ListAgentTickets(ctx context.Context,
	req *gen.ListAgentTicketsRequest) (*gen.ListAgentTicketsResponse, error) {

	rep := repo.NewTicketRepo(s.DB)
	ticketSvc := svc.NewTicketSvc(rep)

	attRep := repo.NewAttachmentRepo(s.DB)
	attSvc := svc.NewAttachmentSvc(attRep)

	ticketViews := ticketSvc.ListAgentTickets(req.TicketStatus)
	ticketsListResp := &gen.ListAgentTicketsResponse{}

	for _, t := range ticketViews {
		ticket := &gen.Ticket{
			Id:                 t.ID.String(),
			CreatedAt:          timestamppb.New(t.CreatedAt),
			Subject:            t.Subject,
			Request:            t.Request,
			Answer:             t.Answer,
			RequesterID:        t.RequesterID.String(),
			Status:             string(t.Status),
			Type:               string(t.TicketType),
			Priority:           string(t.Priority),
			Rate:               uint32(t.Rate),
			RequesterEmail:     t.Email,
			RequesterFullName:  t.FirstName + " " + t.LastName,
			RequestAttachments: attSvc.TicketAttachment(t.ID.String(), axone.ATTACHMENT_KIND_REQUEST),
			AnswerAttachments:  attSvc.TicketAttachment(t.ID.String(), axone.ATTACHMENT_KIND_ANSWER),
		}

		ticketsListResp.Tickets = append(ticketsListResp.Tickets, ticket)

	}

	return ticketsListResp, nil
}

//Subscribe store stream to notify agents when a new request arrive
func (s *AxoneServer) Subscribe(req *gen.NotificationRequest, stream gen.Axone_SubscribeServer) error {
	fin := make(chan bool)
	s.subscribers.Store(req.Id, sub{stream: stream, finished: fin})

	ctx := stream.Context()
	// Keep this scope alive because once this scope exits - the stream is closed
	for {
		select {
		case <-fin:
			log.Printf("Closing stream for client ID: %s", req.Id)
			return nil
		case <-ctx.Done():
			log.Printf("Client ID %s has disconnected", req.Id)
			return nil
		}
	}

}

func (s *AxoneServer) Unsubscribe(ctx context.Context, req *gen.NotificationRequest) (*gen.NotificationResponse, error) {
	v, ok := s.subscribers.Load(req.Id)
	if !ok {
		return nil, fmt.Errorf("failed to load subscriber key: %s", req.Id)
	}
	sub, ok := v.(sub)
	if !ok {
		return nil, fmt.Errorf("failed to cast subscriber value: %T", v)
	}
	select {
	case sub.finished <- true:
		log.Printf("Unsubscribed client: %s", req.Id)
	default:
		// Default case is to avoid blocking in case client has already unsubscribed
	}
	s.subscribers.Delete(req.Id)
	return &gen.NotificationResponse{}, nil
}

//SendNotification notify agents when new tickets arrive
func (s *AxoneServer) SendNotification(msg string) {

	// A list of clients to unsubscribe in case of error
	var unsubscribe []string

	s.subscribers.Range(func(k, v interface{}) bool {
		id, ok := k.(string)
		if !ok {
			log.Printf("Failed to cast subscriber key: %T", k)
			return false
		}
		sub, ok := v.(sub)
		if !ok {
			log.Printf("Failed to cast subscriber value: %T", v)
			return false
		}
		notification := &gen.NotificationResponse{
			Message: msg,
			Time:    time.Now().UnixNano(),
		}
		if err := sub.stream.Send(notification); err != nil {
			log.Printf("Failed to send data to client: %v", err)
			select {
			case sub.finished <- true:
				log.Printf("Unsubscribed client: %s", id)
			default:
				// Default case is to avoid blocking in case client has already unsubscribed
			}
			// In case of error the client would re-subscribe so close the subscriber stream
			unsubscribe = append(unsubscribe, id)
		}
		return true
	})
	// Unsubscribe erroneous client streams
	for _, id := range unsubscribe {
		log.Printf("Unsubscribing client: %s", id)
		s.subscribers.Delete(id)
	}
}

func (s *AxoneServer) ListRequesterTickets(ctx context.Context,
	req *gen.ListRequesterTicketsRequest) (*gen.ListRequesterTicketsResponse, error) {

	rep := repo.NewTicketRepo(s.DB)
	ticketSvc := svc.NewTicketSvc(rep)

	ticketViews := ticketSvc.ListRequesterTickets(req.TicketStatus, req.RequesterID)
	ticketsListResp := &gen.ListRequesterTicketsResponse{}

	for _, t := range ticketViews {
		ticket := &gen.Ticket{
			Subject:           t.Subject,
			Request:           t.Request,
			Answer:            t.Answer,
			RequesterID:       t.RequesterID.String(),
			Status:            string(t.Status),
			Type:              string(t.TicketType),
			Priority:          string(t.Priority),
			Rate:              uint32(t.Rate),
			RequesterLogin:    t.Login,
			RequesterEmail:    t.Email,
			RequesterFullName: t.FirstName + " " + t.LastName,
			CreatedAt:         timestamppb.New(t.CreatedAt),
			UpdatedAt:         timestamppb.New(t.UpdatedAt),
		}

		ticketsListResp.Tickets = append(ticketsListResp.Tickets, ticket)

	}

	return ticketsListResp, nil

}

func (s *AxoneServer) ListTags(ctx context.Context, req *gen.ListTagRequest) (*gen.ListTagResponse, error) {
	repo := repo.NewTagRepo(s.DB)
	tagSvc := svc.NewTagSvc(repo)

	tags, err := tagSvc.ListTags()
	if err != nil {
		return &gen.ListTagResponse{}, err
	}

	tagsResp := &gen.ListTagResponse{}

	for _, t := range tags {
		tag := &gen.Tag{
			Id:    t.ID.String(),
			Key:   t.Key,
			Value: t.Value,
		}
		tagsResp.Tags = append(tagsResp.Tags, tag)
	}
	return tagsResp, nil
}

func (s *AxoneServer) AddTag(ctx context.Context, req *gen.AddTagRequest) (*gen.AddTagResponse, error) {
	repo := repo.NewTicketRepo(s.DB)
	ticketSvc := svc.NewTicketSvc(repo)

	err := ticketSvc.AddTag(req.TicketID, req.TagID)
	if err != nil {
		return &gen.AddTagResponse{}, err
	}
	return &gen.AddTagResponse{}, nil
}
