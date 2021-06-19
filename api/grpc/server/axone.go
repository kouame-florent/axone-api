package server

import (
	"bytes"
	"context"
	"io"
	"log"

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
)

const maxAttchmentSize = 20 << 20 //1 Mi

type AxoneServer struct {
	gen.UnimplementedAxoneServer
	TicketSvc *svc.TicketSvc
}

func NewAxoneServer(svc *svc.TicketSvc) *AxoneServer {
	return &AxoneServer{
		TicketSvc: svc,
	}
}

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

	id, err := s.TicketSvc.SendNewTicket(ticketID, req.Subject, req.Request, axone.TicketType(req.Type), requesterID)
	if err != nil {
		return &gen.NewTicketResponse{}, err
	}

	resp := &gen.NewTicketResponse{
		ID: id.String(),
	}

	return resp, nil
}

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

	db := store.NewDB()
	rep := repo.NewAttachmentRepo(db)
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

	res := &gen.AttachmentResponse{
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
	req *gen.AgentTicketsListRequest) (*gen.AgentTicketsListResponse, error) {

	db := store.NewDB()
	rep := repo.NewTicketRepo(db)
	ticketSvc := svc.NewTicketSvc(rep)

	ticketViews := ticketSvc.ListAgentTickets(req.Status)
	ticketsListResp := &gen.AgentTicketsListResponse{}

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
			RequesterEmail:    t.Email,
			RequesterFullName: t.FirstName + " " + t.LastName,
		}

		ticketsListResp.Tickets = append(ticketsListResp.Tickets, ticket)

	}

	return ticketsListResp, nil
}

/*
func saveMeta(db *gorm.DB, req *gen.AttachmentRequest, storgeName string) (uuid.UUID, error) {
	meta := req.GetInfo()
	storageName := storgeName
	tickerID, err := uuid.Parse(meta.GetTicketID())
	if err != nil {
		return uuid.UUID{}, err
	}

	att := &axone.Attachment{
		Model: axone.Model{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		UploadedName: meta.GetUploadedName(),
		MimeType:     meta.GetMimeType(),
		Size:         meta.GetSize(),
		StorageName:  storageName,
		Kind:         axone.ATTACHMENT_KIND_REQUEST,
		TicketID:     tickerID,
	}

	rep := repo.NewAttachmentRepo(db)
	return rep.Create(att)

}
*/
