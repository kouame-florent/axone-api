package svc

import (
	"time"

	"github.com/google/uuid"
	"github.com/kouame-florent/axone-api/internal/axone"
	"github.com/kouame-florent/axone-api/internal/config"
	"github.com/kouame-florent/axone-api/internal/repo"
	"github.com/kouame-florent/axone-api/internal/store"
)

func Createfakeorganization() (uuid.UUID, error) {
	dsn := config.DataSourceName()
	db := store.OpenDB(dsn)

	repo := repo.NewOrganizationRepo(db)

	o := &axone.Organization{
		Model: axone.Model{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name: "DTI",
	}
	err := repo.DB.Create(o).Error
	if err != nil {
		return o.ID, err
	}
	return uuid.UUID{}, nil
}

func CreateFakeUser(organizationID uuid.UUID) (uuid.UUID, error) {
	dsn := config.DataSourceName()
	db := store.OpenDB(dsn)

	repo := repo.NewUserRepo(db)

	u := &axone.User{
		Model: axone.Model{
			ID:        uuid.New(),
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

	err := repo.DB.Create(u).Error
	if err != nil {
		return u.ID, err
	}
	return uuid.UUID{}, nil
}

func CreateEndUser(userID uuid.UUID) (uuid.UUID, error) {
	dsn := config.DataSourceName()
	db := store.OpenDB(dsn)

	repo := repo.NewUserRepo(db)

	eu := &axone.EndUser{
		UserID: userID,
	}
	err := repo.DB.Create(eu).Error
	if err != nil {
		return userID, err
	}
	return uuid.UUID{}, nil
}
