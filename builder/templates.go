package main

var repoTpl = `
package repo

import (
	"errors"

	"github.com/google/uuid"
	"github.com/kouame-florent/icens-api/internal/icens"
	"gorm.io/gorm"
)

type {{ .Entity }}Repo struct {
	DB *gorm.DB
}

func New{{ .Entity }}Repo(db *gorm.DB) *{{ .Entity }}Repo {
	return &{{ .Entity }}Repo{
		DB: db,
	}
}

func (r *{{ .Entity }}Repo) Create(e *icens.{{ .Entity }}) (uuid.UUID, error) {
	tx := r.DB.Create(e)
	if tx.Error != nil {
		return uuid.UUID{}, tx.Error
	}
	return e.ID, nil
}

func (r *{{ .Entity }}Repo) Find(id uuid.UUID) (*icens.{{ .Entity }}, error) {
	e := &icens.{{ .Entity }}{}
	tx := r.DB.First(e, id)
	if tx.Error != nil {
		return &icens.{{ .Entity }}{}, tx.Error
	}
	
	return e, nil
}

func (r *{{ .Entity }}Repo) FindAll() ([]icens.{{ .Entity }}, error) {
	var ents []icens.{{ .Entity }}
	tx := r.DB.Find(&ents)
	if tx.Error != nil {
		return []icens.{{ .Entity }}{}, tx.Error
	}
	return ents, nil
}

func (r *{{ .Entity }}Repo) FindRange(offset, size int) ([]icens.{{ .Entity }}, error) {
	return []icens.{{ .Entity }}{}, nil
}

func (r *{{ .Entity }}Repo) Count() int {
	return 0
}

func (r *{{ .Entity }}Repo) Update(l *icens.{{ .Entity }}) error {
	return nil
}

func (r *{{ .Entity }}Repo) Delete(l *icens.{{ .Entity }}) error {
	return nil
}

`
