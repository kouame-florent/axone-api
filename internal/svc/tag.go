package svc

import (
	"time"

	"github.com/google/uuid"
	"github.com/kouame-florent/axone-api/internal/axone"
	"github.com/kouame-florent/axone-api/internal/repo"
)

type TagSvc struct {
	repo *repo.TagRepo
}

func NewTagSvc(rep *repo.TagRepo) *TagSvc {
	return &TagSvc{
		repo: rep,
	}
}

func (s *TagSvc) ListTags() ([]axone.Tag, error) {
	return s.repo.FindAll()
}

func (s *TagSvc) Create(status axone.TagStatus, key, value, description string) (uuid.UUID, error) {
	t := &axone.Tag{
		Model: axone.Model{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Key:         key,
		Value:       value,
		Description: description,
		Status:      status,
	}
	return s.repo.Create(t)
}
