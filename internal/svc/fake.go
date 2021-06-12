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

func (s *FakeSvc) CreateFakeUser(organizationID uuid.UUID) (uuid.UUID, error) {

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
		OrganizationID: organizationID,
	}

	err = repo.DB.Create(u).Error
	if err != nil {

		return uuid.UUID{}, err
	}
	return id, nil
}
