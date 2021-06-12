package repo

import (
	"github.com/google/uuid"
	"github.com/kouame-florent/axone-api/internal/axone"
	"gorm.io/gorm"
)

type RequesterRepo struct {
	DB *gorm.DB
}

func NewRequesterRepo(db *gorm.DB) *RequesterRepo {
	return &RequesterRepo{
		DB: db,
	}
}

func (r *RequesterRepo) Create(e *axone.Requester) (uuid.UUID, error) {
	tx := r.DB.Create(e)
	if tx.Error != nil {
		return uuid.UUID{}, tx.Error
	}
	return e.UserID, nil
}

func (r *RequesterRepo) Find(id uuid.UUID) (*axone.Requester, error) {
	e := &axone.Requester{}
	tx := r.DB.First(e, id)
	if tx.Error != nil {
		return &axone.Requester{}, tx.Error
	}

	return e, nil
}

func (r *RequesterRepo) FindAll() ([]axone.Requester, error) {
	var ents []axone.Requester
	tx := r.DB.Find(&ents)
	if tx.Error != nil {
		return []axone.Requester{}, tx.Error
	}
	return ents, nil
}

func (r *RequesterRepo) FindRange(offset, size int) ([]axone.Requester, error) {
	return []axone.Requester{}, nil
}

func (r *RequesterRepo) Count() int {
	return 0
}

func (r *RequesterRepo) Update(l *axone.Requester) error {
	return nil
}

func (r *RequesterRepo) Delete(l *axone.Requester) error {
	return nil
}
