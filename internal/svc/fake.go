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
	id, err := uuid.Parse("44535ea6-d21d-47bb-8b6e-08e49a4caf64")
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
	id, err := uuid.Parse("8616e9d0-8e56-40b3-b06a-90edd318b7a1")
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
