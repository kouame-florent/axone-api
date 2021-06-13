package server

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/kouame-florent/axone-api/api/grpc/gen"
	"github.com/kouame-florent/axone-api/internal/axone"
	"github.com/kouame-florent/axone-api/internal/svc"
)

type AxoneServer struct {
	gen.UnimplementedAxoneServer
	TicketSvc *svc.TicketSvc
}

func NewAxoneServer(svc *svc.TicketSvc) *AxoneServer {
	return &AxoneServer{
		TicketSvc: svc,
	}
}

func (s *AxoneServer) SendNewTicket(ctx context.Context, nt *gen.NewTicketRequest) (*gen.NewTicketResponse, error) {
	log.Printf("TICKETID: %s", nt.TicketID)
	ticketID, err := uuid.Parse(nt.TicketID)
	if err != nil {
		return &gen.NewTicketResponse{}, err
	}
	requesterID, err := uuid.Parse(nt.RequesterID)
	if err != nil {
		return &gen.NewTicketResponse{}, err
	}

	id, err := s.TicketSvc.SendNewTicket(ticketID, nt.Subject, nt.Request, axone.TicketType(nt.Type), requesterID)
	if err != nil {
		return &gen.NewTicketResponse{}, err
	}

	resp := &gen.NewTicketResponse{
		ID: id.String(),
	}

	return resp, nil
}
