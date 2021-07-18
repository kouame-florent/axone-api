package svc

import (
	"github.com/kouame-florent/axone-api/internal/axone"
	"github.com/kouame-florent/axone-api/internal/repo"
)

type RoleSvc struct {
	Repo *repo.RoleRepo
}

func NewRoleSvc(r *repo.RoleRepo) *RoleSvc {
	return &RoleSvc{
		Repo: r,
	}
}

func (s *RoleSvc) FindByValue(value axone.RoleValue) (*axone.Role, error) {
	var role axone.Role

	err := s.Repo.DB.Where("value = ?", value).First(&role).Error
	if err != nil {
		return &axone.Role{}, err
	}

	return &role, nil

}
