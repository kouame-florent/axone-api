package svc

import (
	"errors"

	"github.com/kouame-florent/axone-api/internal/axone"
	"github.com/kouame-florent/axone-api/internal/repo"
	"gorm.io/gorm"
)

type UserTagsSvc struct {
	DB *gorm.DB
}

func NewUserTagsSvc(db *gorm.DB) *UserTagsSvc {
	return &UserTagsSvc{
		DB: db,
	}
}

func (s *UserTagsSvc) Add(userID string, tagIDs []string) error {
	userRepo := repo.NewUserRepo(s.DB)
	tagRepo := repo.NewTagRepo(s.DB)

	var user axone.User
	err := userRepo.DB.First(&user, "id = ?", userID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	for _, id := range tagIDs {
		var tag axone.Tag
		err = tagRepo.DB.First(&tag, "id = ?", id).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		user.Tags = append(user.Tags, &tag)
	}

	tx := s.DB.Save(&user)
	if tx.Error != nil {
		return err
	}
	return nil
}
