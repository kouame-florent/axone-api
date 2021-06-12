package server

import (
	"context"

	"github.com/google/uuid"
	"github.com/kouame-florent/axone-api/api/grpc/gen"
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

func (s *AxoneServer) SendNewTicket(ctx context.Context, tr *gen.NewTicketRequest) (*gen.NewTicketResponse, error) {
	ticketID, err := uuid.Parse(tr.TicketID)
	if err != nil {
		return &gen.NewTicketResponse{}, err
	}
	requesterID, err := uuid.Parse(tr.RequesterID)
	if err != nil {
		return &gen.NewTicketResponse{}, err
	}

	id, err := s.TicketSvc.SendNewTicket(ticketID, tr.Subject, tr.Request, requesterID)
	if err != nil {
		return &gen.NewTicketResponse{}, err
	}

	resp := &gen.NewTicketResponse{
		ID: id.String(),
	}

	return resp, nil
}
