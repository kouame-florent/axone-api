package svc

import (
	"github.com/google/uuid"
	"github.com/kouame-florent/axone-api/internal/axone"
	"github.com/kouame-florent/axone-api/internal/repo"
)

type AgentSvc struct {
	Repo *repo.AgentRepo
}

func NewAgentSvc(r *repo.AgentRepo) *AgentSvc {
	return &AgentSvc{
		Repo: r,
	}
}

func (s *AgentSvc) CreateAgent(userID uuid.UUID, level axone.AgentLevel, bio string) (uuid.UUID, error) {
	ag := &axone.Agent{
		UserID: userID,
		Level:  level,
		Bio:    bio,
	}
	id, err := s.Repo.Create(ag)
	if err != nil {
		return id, err
	}
	return id, nil
}
