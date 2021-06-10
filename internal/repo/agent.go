package repo

import (
	"github.com/google/uuid"
	"github.com/kouame-florent/axone-api/internal/axone"
	"gorm.io/gorm"
)

type AgentRepo struct {
	DB *gorm.DB
}

func NewAgentRepo(db *gorm.DB) *AgentRepo {
	return &AgentRepo{
		DB: db,
	}
}

func (r *AgentRepo) Create(e *axone.Agent) (uuid.UUID, error) {
	tx := r.DB.Create(e)
	if tx.Error != nil {
		return uuid.UUID{}, tx.Error
	}
	return e.UserID, nil
}

func (r *AgentRepo) Find(id uuid.UUID) (*axone.Agent, error) {
	e := &axone.Agent{}
	tx := r.DB.First(e, id)
	if tx.Error != nil {
		return &axone.Agent{}, tx.Error
	}

	return e, nil
}

func (r *AgentRepo) FindAll() ([]axone.Agent, error) {
	var ents []axone.Agent
	tx := r.DB.Find(&ents)
	if tx.Error != nil {
		return []axone.Agent{}, tx.Error
	}
	return ents, nil
}

func (r *AgentRepo) FindRange(offset, size int) ([]axone.Agent, error) {
	return []axone.Agent{}, nil
}

func (r *AgentRepo) Count() int {
	return 0
}

func (r *AgentRepo) Update(l *axone.Agent) error {
	return nil
}

func (r *AgentRepo) Delete(l *axone.Agent) error {
	return nil
}
