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
		TicketType:  axone.TicketType(ticketType),
		RequesterID: requesterID,
		Status:      axone.TICKET_STATUS_NEW,
	}

	id, err := s.Repo.Create(t)
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil

}

/*
func (s *TicketSvc) SendAttachment() (uuid.UUID, error) {
	return uuid.UUID{}, nil
}
*/

type ListTicketsResult struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Subject     string
	Request     string
	Answer      string
	RequesterID uuid.UUID
	Email       string
	FirstName   string
	LastName    string
	Status      string
	TicketType  string
	Priority    string
	Rate        uint
}

func (s *TicketSvc) ListAgentTickets(status string) []ListTicketsResult {
	//var tickets []axone.Ticket
	var results []ListTicketsResult
	s.Repo.DB.Model(&axone.Ticket{}).Select("tickets.id,tickets.created_at,tickets.updated_at,tickets.subject, tickets.request,tickets.answer,tickets.requester_id,tickets.status,tickets.ticket_type,tickets.priority,tickets.rate,users.email,users.first_name,users.last_name").
		Joins("join requesters on requesters.user_id = tickets.requester_id").Joins("join users on requesters.user_id = users.id").
		Where("status = ?", status).Scan(&results)

	return results
}

/*
func toTickets(results []listTicketsResult) []axone.Ticket {
	var tickets []axone.Ticket

	for _, res := range results {
		tk := axone.Ticket{
			Model: axone.Model{
				ID:        res.ID,
				CreatedAt: res.CreatedAt,
				UpdatedAt: res.UpdatedAt,
			},
			Subject:     res.Subject,
			Request:     res.Request,
			Answer:      res.Answer,
			RequesterID: res.RequesterID,
			Status:      axone.TicketStatus(res.Status),
			TicketType:  axone.TicketType(res.TicketType),
			Priority:    axone.TicketPriority(res.Priority),
			Rate:        uint(res.Rate),
		}
		tickets = append(tickets, tk)

	}

	return tickets
}
*/
