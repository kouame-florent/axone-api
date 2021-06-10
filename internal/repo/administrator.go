package repo

import (
	"github.com/google/uuid"
	"github.com/kouame-florent/axone-api/internal/axone"
	"gorm.io/gorm"
)

type AdministratorRepo struct {
	DB *gorm.DB
}

func NewAdministratorRepo(db *gorm.DB) *AdministratorRepo {
	return &AdministratorRepo{
		DB: db,
	}
}

func (r *AdministratorRepo) Create(e *axone.Administrator) (uuid.UUID, error) {
	tx := r.DB.Create(e)
	if tx.Error != nil {
		return uuid.UUID{}, tx.Error
	}
	return e.UserID, nil
}

func (r *AdministratorRepo) Find(id uuid.UUID) (*axone.Administrator, error) {
	e := &axone.Administrator{}
	tx := r.DB.First(e, id)
	if tx.Error != nil {
		return &axone.Administrator{}, tx.Error
	}

	return e, nil
}

func (r *AdministratorRepo) FindAll() ([]axone.Administrator, error) {
	var ents []axone.Administrator
	tx := r.DB.Find(&ents)
	if tx.Error != nil {
		return []axone.Administrator{}, tx.Error
	}
	return ents, nil
}

func (r *AdministratorRepo) FindRange(offset, size int) ([]axone.Administrator, error) {
	return []axone.Administrator{}, nil
}

func (r *AdministratorRepo) Count() int {
	return 0
}

func (r *AdministratorRepo) Update(l *axone.Administrator) error {
	return nil
}

func (r *AdministratorRepo) Delete(l *axone.Administrator) error {
	return nil
}
