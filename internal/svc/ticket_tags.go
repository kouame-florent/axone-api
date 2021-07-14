package svc

import (
	"github.com/kouame-florent/axone-api/internal/axone"
	"gorm.io/gorm"
)

type TicketTagsSvc struct {
	DB *gorm.DB
}

func NewTicketTagsSvc(db *gorm.DB) *TicketTagsSvc {
	return &TicketTagsSvc{
		DB: db,
	}
}

func (t *TicketTagsSvc) Exist(ticketID, tagID string) (bool, error) {

	res := axone.TicketTags{}

	if tx := t.DB.Table("ticket_tags").Where("ticket_id = ? AND tag_id = ?", ticketID, tagID).Scan(&res); tx.Error != nil {
		return false, tx.Error
	}

	if (axone.TicketTags{}) == res {
		return false, nil
	}

	return true, nil

}

func (t *TicketTagsSvc) Remove(ticket *axone.Ticket, tag *axone.Tag) error {
	return t.DB.Model(ticket).Association("Tags").Delete(tag)

	//return nil
}
