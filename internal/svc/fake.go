package svc

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/kouame-florent/axone-api/internal/axone"
	"github.com/kouame-florent/axone-api/internal/repo"
	"gorm.io/gorm"
)

type FakeSvc struct {
	DB *gorm.DB
}

func NewFakeSvc(db *gorm.DB) *FakeSvc {
	return &FakeSvc{
		DB: db,
	}
}

func (s *FakeSvc) CreatefakeOrganization() (uuid.UUID, error) {

	repo := repo.NewOrganizationRepo(s.DB)
	id, err := uuid.Parse("fc16897d-82eb-45c3-b0c6-31bb71cf391b")
	if err != nil {
		log.Fatal(err)
	}
	o := &axone.Organization{
		Model: axone.Model{
			ID:        id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name: "Central nucléaire",
	}
	err = repo.DB.Create(o).Error
	if err != nil {
		return uuid.UUID{}, err
	}
	return id, nil
}

func (s *FakeSvc) CreateFakeRequesterUser(organizationID uuid.UUID) (uuid.UUID, error) {

	repo := repo.NewUserRepo(s.DB)
	id, err := uuid.Parse("4a2bfb72-94ab-4fb2-b195-52dc1a12ffdb")
	if err != nil {
		log.Fatal(err)
	}
	u := &axone.User{
		Model: axone.Model{
			ID:        id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		FirstName:      "Homer",
		LastName:       "Simpson",
		Email:          "homer.simpson@gmail.com",
		PhoneNumber:    "225-05-05-45-78-65",
		Login:          "homer",
		Password:       "homer",
		Status:         axone.USER_STATUS_ENABLED,
		OrganizationID: organizationID,
	}

	err = repo.DB.Create(u).Error
	if err != nil {

		return uuid.UUID{}, err
	}
	return id, nil
}

func (s *FakeSvc) CreateFakeLevelOneAgentUser(organizationID uuid.UUID) (uuid.UUID, error) {

	repo := repo.NewUserRepo(s.DB)
	id, err := uuid.Parse("56a680d0-47c4-48b8-9ad5-6eff936d4d75")
	if err != nil {
		log.Fatal(err)
	}
	u := &axone.User{
		Model: axone.Model{
			ID:        id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		FirstName:      "Lisa",
		LastName:       "Simpson",
		Email:          "lisa.simpson@gmail.com",
		PhoneNumber:    "225-04-04-44-78-65",
		Login:          "lisa",
		Password:       "lisa",
		Status:         axone.USER_STATUS_ENABLED,
		OrganizationID: organizationID,
	}

	err = repo.DB.Create(u).Error
	if err != nil {

		return uuid.UUID{}, err
	}
	return id, nil
}
