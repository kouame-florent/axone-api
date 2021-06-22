package svc

import (
	"github.com/google/uuid"
	"github.com/kouame-florent/axone-api/internal/axone"
	"github.com/kouame-florent/axone-api/internal/repo"
)

type RequesterSvc struct {
	Repo *repo.RequesterRepo
}

func NewRequesterSvc(r *repo.RequesterRepo) *RequesterSvc {
	return &RequesterSvc{
		Repo: r,
	}
}

func (s *RequesterSvc) CreateRequester(userID uuid.UUID) (uuid.UUID, error) {
	eu := &axone.Requester{
		UserID: userID,
	}
	id, err := s.Repo.Create(eu)
	if err != nil {
		return id, err
	}
	return id, nil
}
