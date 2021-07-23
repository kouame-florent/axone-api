package svc

import (
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/kouame-florent/axone-api/internal/axone"
	"github.com/kouame-florent/axone-api/internal/repo"
	"gorm.io/gorm"
)

type Initialization struct {
	RoleRepo *repo.RoleRepo
}

func NewInitialization(roleRepo *repo.RoleRepo) *Initialization {
	return &Initialization{
		RoleRepo: roleRepo,
	}
}

func (i *Initialization) CreateDefaultRoles() {
	repo := repo.NewRoleRepo(i.RoleRepo.DB)
	svc := NewRoleSvc(repo)

	agentRole := axone.Role{
		Model: axone.Model{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Value: axone.ROLE_VALUE_AGENT,
	}

	admintRole := axone.Role{
		Model: axone.Model{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Value: axone.ROLE_VALUE_ADMINISTRATOR,
	}

	requesterRole := axone.Role{
		Model: axone.Model{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Value: axone.ROLE_VALUE_REQUESTER,
	}

	//check if not exist
	_, err := svc.FindByValue(axone.ROLE_VALUE_AGENT)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			svc.Repo.Create(&agentRole)
		} else {
			log.Print(err)
			panic(err)
		}
	}

	_, err = svc.FindByValue(axone.ROLE_VALUE_ADMINISTRATOR)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			svc.Repo.Create(&admintRole)
		} else {
			log.Print(err)
			panic(err)
		}

	}

	_, err = svc.FindByValue(axone.ROLE_VALUE_REQUESTER)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			svc.Repo.Create(&requesterRole)
		} else {
			log.Print(err)
			panic(err)
		}
	}

}
