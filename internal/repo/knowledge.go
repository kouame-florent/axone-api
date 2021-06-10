
package repo

import (

	"github.com/google/uuid"
	"github.com/kouame-florent/axone-api/internal/axone"
	"gorm.io/gorm"
)

type KnowledgeRepo struct {
	DB *gorm.DB
}

func NewKnowledgeRepo(db *gorm.DB) *KnowledgeRepo {
	return &KnowledgeRepo{
		DB: db,
	}
}

func (r *KnowledgeRepo) Create(e *axone.Knowledge) (uuid.UUID, error) {
	tx := r.DB.Create(e)
	if tx.Error != nil {
		return uuid.UUID{}, tx.Error
	}
	return e.ID, nil
}

func (r *KnowledgeRepo) Find(id uuid.UUID) (*axone.Knowledge, error) {
	e := &axone.Knowledge{}
	tx := r.DB.First(e, id)
	if tx.Error != nil {
		return &axone.Knowledge{}, tx.Error
	}
	
	return e, nil
}

func (r *KnowledgeRepo) FindAll() ([]axone.Knowledge, error) {
	var ents []axone.Knowledge
	tx := r.DB.Find(&ents)
	if tx.Error != nil {
		return []axone.Knowledge{}, tx.Error
	}
	return ents, nil
}

func (r *KnowledgeRepo) FindRange(offset, size int) ([]axone.Knowledge, error) {
	return []axone.Knowledge{}, nil
}

func (r *KnowledgeRepo) Count() int {
	return 0
}

func (r *KnowledgeRepo) Update(l *axone.Knowledge) error {
	return nil
}

func (r *KnowledgeRepo) Delete(l *axone.Knowledge) error {
	return nil
}

