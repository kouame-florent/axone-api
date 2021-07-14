package svc

import (
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

func (t *TagSvc) ListTags() ([]axone.Tag, error) {
	return t.repo.FindAll()
}
