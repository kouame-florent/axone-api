package main

var repoTpl = `
package repo

import (

	"github.com/google/uuid"
	"github.com/kouame-florent/axone-api/internal/axone"
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

func (r *{{ .Entity }}Repo) Create(e *axone.{{ .Entity }}) (uuid.UUID, error) {
	tx := r.DB.Create(e)
	if tx.Error != nil {
		return uuid.UUID{}, tx.Error
	}
	return e.ID, nil
}

func (r *{{ .Entity }}Repo) Find(id uuid.UUID) (*axone.{{ .Entity }}, error) {
	e := &axone.{{ .Entity }}{}
	tx := r.DB.First(e, id)
	if tx.Error != nil {
		return &axone.{{ .Entity }}{}, tx.Error
	}
	
	return e, nil
}

func (r *{{ .Entity }}Repo) FindAll() ([]axone.{{ .Entity }}, error) {
	var ents []axone.{{ .Entity }}
	tx := r.DB.Find(&ents)
	if tx.Error != nil {
		return []axone.{{ .Entity }}{}, tx.Error
	}
	return ents, nil
}

func (r *{{ .Entity }}Repo) FindRange(offset, size int) ([]axone.{{ .Entity }}, error) {
	return []axone.{{ .Entity }}{}, nil
}

func (r *{{ .Entity }}Repo) Count() int {
	return 0
}

func (r *{{ .Entity }}Repo) Update(l *axone.{{ .Entity }}) error {
	return nil
}

func (r *{{ .Entity }}Repo) Delete(l *axone.{{ .Entity }}) error {
	return nil
}

`
