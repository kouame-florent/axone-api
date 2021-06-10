
package repo

import (

	"github.com/google/uuid"
	"github.com/kouame-florent/axone-api/internal/axone"
	"gorm.io/gorm"
)

type RoleRepo struct {
	DB *gorm.DB
}

func NewRoleRepo(db *gorm.DB) *RoleRepo {
	return &RoleRepo{
		DB: db,
	}
}

func (r *RoleRepo) Create(e *axone.Role) (uuid.UUID, error) {
	tx := r.DB.Create(e)
	if tx.Error != nil {
		return uuid.UUID{}, tx.Error
	}
	return e.ID, nil
}

func (r *RoleRepo) Find(id uuid.UUID) (*axone.Role, error) {
	e := &axone.Role{}
	tx := r.DB.First(e, id)
	if tx.Error != nil {
		return &axone.Role{}, tx.Error
	}
	
	return e, nil
}

func (r *RoleRepo) FindAll() ([]axone.Role, error) {
	var ents []axone.Role
	tx := r.DB.Find(&ents)
	if tx.Error != nil {
		return []axone.Role{}, tx.Error
	}
	return ents, nil
}

func (r *RoleRepo) FindRange(offset, size int) ([]axone.Role, error) {
	return []axone.Role{}, nil
}

func (r *RoleRepo) Count() int {
	return 0
}

func (r *RoleRepo) Update(l *axone.Role) error {
	return nil
}

func (r *RoleRepo) Delete(l *axone.Role) error {
	return nil
}

