
package repo

import (

	"github.com/google/uuid"
	"github.com/kouame-florent/axone-api/internal/axone"
	"gorm.io/gorm"
)

type AssignmentRepo struct {
	DB *gorm.DB
}

func NewAssignmentRepo(db *gorm.DB) *AssignmentRepo {
	return &AssignmentRepo{
		DB: db,
	}
}

func (r *AssignmentRepo) Create(e *axone.Assignment) (uuid.UUID, error) {
	tx := r.DB.Create(e)
	if tx.Error != nil {
		return uuid.UUID{}, tx.Error
	}
	return e.ID, nil
}

func (r *AssignmentRepo) Find(id uuid.UUID) (*axone.Assignment, error) {
	e := &axone.Assignment{}
	tx := r.DB.First(e, id)
	if tx.Error != nil {
		return &axone.Assignment{}, tx.Error
	}
	
	return e, nil
}

func (r *AssignmentRepo) FindAll() ([]axone.Assignment, error) {
	var ents []axone.Assignment
	tx := r.DB.Find(&ents)
	if tx.Error != nil {
		return []axone.Assignment{}, tx.Error
	}
	return ents, nil
}

func (r *AssignmentRepo) FindRange(offset, size int) ([]axone.Assignment, error) {
	return []axone.Assignment{}, nil
}

func (r *AssignmentRepo) Count() int {
	return 0
}

func (r *AssignmentRepo) Update(l *axone.Assignment) error {
	return nil
}

func (r *AssignmentRepo) Delete(l *axone.Assignment) error {
	return nil
}

