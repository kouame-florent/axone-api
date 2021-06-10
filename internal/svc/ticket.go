package svc

import (
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

func (s *TicketSvc) SendTicket() {

}
