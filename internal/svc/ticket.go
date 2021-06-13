package svc

import (
	"time"

	"github.com/google/uuid"
	"github.com/kouame-florent/axone-api/internal/axone"
	"github.com/kouame-florent/axone-api/internal/repo"
)

type TicketSvc struct {
	Repo *repo.TicketRepo
}

func NewTicketSvc(r *repo.TicketRepo) *TicketSvc {
	return &TicketSvc{
		Repo: r,
	}
}

//send ticket and return it ID
func (s *TicketSvc) SendNewTicket(ticketID uuid.UUID, subject, request string, ticketType axone.TicketType, requesterID uuid.UUID) (uuid.UUID, error) {
	t := &axone.Ticket{
		Model: axone.Model{
			ID:        ticketID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Subject:     subject,
		Request:     request,
		Type:        axone.TicketType(ticketType),
		RequesterID: requesterID,
		Status:      axone.TICKET_STATUS_NEW,
	}

	id, err := s.Repo.Create(t)
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil

}

func (s *TicketSvc) SendAttachment() (uuid.UUID, error) {
	return uuid.UUID{}, nil
}
