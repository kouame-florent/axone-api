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

func (s *FakeSvc) createFakeTags() ([]uuid.UUID, error) {

	repo := repo.NewTagRepo(s.DB)

	tags := []axone.Tag{
		{
			Model: axone.Model{
				ID:        uuid.MustParse("99fbaff3-321d-4829-86d1-70a47c6ec020"),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Key:    "service",
			Value:  "security",
			Status: axone.TAG_STATUS_PRIVATE,
		},
		{
			Model: axone.Model{
				ID:        uuid.MustParse("d20e955a-7491-4cd4-9e7d-c1434f4b5c43"),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Key:    "service",
			Value:  "finance",
			Status: axone.TAG_STATUS_PRIVATE,
		},
		{
			Model: axone.Model{
				ID:        uuid.MustParse("1b800073-263a-4ff0-ba42-d0fd0a779900"),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Key:    "developper",
			Value:  "Bill Gates",
			Status: axone.TAG_STATUS_PRIVATE,
		},
		{
			Model: axone.Model{
				ID:        uuid.MustParse("2083a0c7-932f-4c1c-94b8-f472b96a2855"),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Key:    "developper",
			Value:  "Linus Torvalds",
			Status: axone.TAG_STATUS_PRIVATE,
		},
		{
			Model: axone.Model{
				ID:        uuid.MustParse("058a6fe7-fe2c-4661-bbfa-9f1b0ea83481"),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Key:    "tester",
			Value:  "Rob Pike",
			Status: axone.TAG_STATUS_PRIVATE,
		},
		{
			Model: axone.Model{
				ID:        uuid.MustParse("4fffb491-7f6c-45a3-bbd5-1298d58ce16c"),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Key:    "tester",
			Value:  "Bill Joy",
			Status: axone.TAG_STATUS_PRIVATE,
		},
	}

	err := repo.DB.Create(tags).Error
	if err != nil {
		return []uuid.UUID{}, nil
	}

	res := []uuid.UUID{}
	for _, t := range tags {
		res = append(res, t.ID)
	}

	return res, nil
}
