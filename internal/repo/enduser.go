package repo

import (
	"github.com/google/uuid"
	"github.com/kouame-florent/axone-api/internal/axone"
	"gorm.io/gorm"
)

type EndUserRepo struct {
	DB *gorm.DB
}

func NewEndUserRepo(db *gorm.DB) *EndUserRepo {
	return &EndUserRepo{
		DB: db,
	}
}

func (r *EndUserRepo) Create(e *axone.EndUser) (uuid.UUID, error) {
	tx := r.DB.Create(e)
	if tx.Error != nil {
		return uuid.UUID{}, tx.Error
	}
	return e.UserID, nil
}

func (r *EndUserRepo) Find(id uuid.UUID) (*axone.EndUser, error) {
	e := &axone.EndUser{}
	tx := r.DB.First(e, id)
	if tx.Error != nil {
		return &axone.EndUser{}, tx.Error
	}

	return e, nil
}

func (r *EndUserRepo) FindAll() ([]axone.EndUser, error) {
	var ents []axone.EndUser
	tx := r.DB.Find(&ents)
	if tx.Error != nil {
		return []axone.EndUser{}, tx.Error
	}
	return ents, nil
}

func (r *EndUserRepo) FindRange(offset, size int) ([]axone.EndUser, error) {
	return []axone.EndUser{}, nil
}

func (r *EndUserRepo) Count() int {
	return 0
}

func (r *EndUserRepo) Update(l *axone.EndUser) error {
	return nil
}

func (r *EndUserRepo) Delete(l *axone.EndUser) error {
	return nil
}
