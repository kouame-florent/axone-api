
package repo

import (

	"github.com/google/uuid"
	"github.com/kouame-florent/axone-api/internal/axone"
	"gorm.io/gorm"
)

type TicketRepo struct {
	DB *gorm.DB
}

func NewTicketRepo(db *gorm.DB) *TicketRepo {
	return &TicketRepo{
		DB: db,
	}
}

func (r *TicketRepo) Create(e *axone.Ticket) (uuid.UUID, error) {
	tx := r.DB.Create(e)
	if tx.Error != nil {
		return uuid.UUID{}, tx.Error
	}
	return e.ID, nil
}

func (r *TicketRepo) Find(id uuid.UUID) (*axone.Ticket, error) {
	e := &axone.Ticket{}
	tx := r.DB.First(e, id)
	if tx.Error != nil {
		return &axone.Ticket{}, tx.Error
	}
	
	return e, nil
}

func (r *TicketRepo) FindAll() ([]axone.Ticket, error) {
	var ents []axone.Ticket
	tx := r.DB.Find(&ents)
	if tx.Error != nil {
		return []axone.Ticket{}, tx.Error
	}
	return ents, nil
}

func (r *TicketRepo) FindRange(offset, size int) ([]axone.Ticket, error) {
	return []axone.Ticket{}, nil
}

func (r *TicketRepo) Count() int {
	return 0
}

func (r *TicketRepo) Update(l *axone.Ticket) error {
	return nil
}

func (r *TicketRepo) Delete(l *axone.Ticket) error {
	return nil
}

