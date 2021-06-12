package svc

import (
	"github.com/google/uuid"
	"github.com/kouame-florent/axone-api/internal/axone"
	"github.com/kouame-florent/axone-api/internal/repo"
)

type RequesterSvc struct {
	Repo *repo.RequesterRepo
}

func NewEndUserSvc(r *repo.RequesterRepo) *RequesterSvc {
	return &RequesterSvc{
		Repo: r,
	}
}

func (s *RequesterSvc) CreateEndUser(userID uuid.UUID) (uuid.UUID, error) {
	eu := &axone.Requester{
		UserID: userID,
	}
	id, err := s.Repo.Create(eu)
	if err != nil {
		return id, err
	}
	return uuid.UUID{}, nil
}
