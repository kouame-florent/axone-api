package svc

import (
	"github.com/google/uuid"
	"github.com/kouame-florent/axone-api/internal/axone"
	"github.com/kouame-florent/axone-api/internal/repo"
)

type AdministratorSvc struct {
	Repo *repo.AdministratorRepo
}

func NewAdministratorSvc(r *repo.AdministratorRepo) *AdministratorSvc {
	return &AdministratorSvc{
		Repo: r,
	}
}

func (s *AdministratorSvc) CreateAdministrator(userID uuid.UUID) (uuid.UUID, error) {
	ad := &axone.Administrator{
		UserID: userID,
	}
	id, err := s.Repo.Create(ad)
	if err != nil {
		return id, err
	}
	return id, nil
}
