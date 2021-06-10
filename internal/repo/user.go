
package repo

import (

	"github.com/google/uuid"
	"github.com/kouame-florent/axone-api/internal/axone"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (r *UserRepo) Create(e *axone.User) (uuid.UUID, error) {
	tx := r.DB.Create(e)
	if tx.Error != nil {
		return uuid.UUID{}, tx.Error
	}
	return e.ID, nil
}

func (r *UserRepo) Find(id uuid.UUID) (*axone.User, error) {
	e := &axone.User{}
	tx := r.DB.First(e, id)
	if tx.Error != nil {
		return &axone.User{}, tx.Error
	}
	
	return e, nil
}

func (r *UserRepo) FindAll() ([]axone.User, error) {
	var ents []axone.User
	tx := r.DB.Find(&ents)
	if tx.Error != nil {
		return []axone.User{}, tx.Error
	}
	return ents, nil
}

func (r *UserRepo) FindRange(offset, size int) ([]axone.User, error) {
	return []axone.User{}, nil
}

func (r *UserRepo) Count() int {
	return 0
}

func (r *UserRepo) Update(l *axone.User) error {
	return nil
}

func (r *UserRepo) Delete(l *axone.User) error {
	return nil
}

