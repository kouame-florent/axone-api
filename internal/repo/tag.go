
package repo

import (

	"github.com/google/uuid"
	"github.com/kouame-florent/axone-api/internal/axone"
	"gorm.io/gorm"
)

type TagRepo struct {
	DB *gorm.DB
}

func NewTagRepo(db *gorm.DB) *TagRepo {
	return &TagRepo{
		DB: db,
	}
}

func (r *TagRepo) Create(e *axone.Tag) (uuid.UUID, error) {
	tx := r.DB.Create(e)
	if tx.Error != nil {
		return uuid.UUID{}, tx.Error
	}
	return e.ID, nil
}

func (r *TagRepo) Find(id uuid.UUID) (*axone.Tag, error) {
	e := &axone.Tag{}
	tx := r.DB.First(e, id)
	if tx.Error != nil {
		return &axone.Tag{}, tx.Error
	}
	
	return e, nil
}

func (r *TagRepo) FindAll() ([]axone.Tag, error) {
	var ents []axone.Tag
	tx := r.DB.Find(&ents)
	if tx.Error != nil {
		return []axone.Tag{}, tx.Error
	}
	return ents, nil
}

func (r *TagRepo) FindRange(offset, size int) ([]axone.Tag, error) {
	return []axone.Tag{}, nil
}

func (r *TagRepo) Count() int {
	return 0
}

func (r *TagRepo) Update(l *axone.Tag) error {
	return nil
}

func (r *TagRepo) Delete(l *axone.Tag) error {
	return nil
}

