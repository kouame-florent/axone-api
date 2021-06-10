
package repo

import (

	"github.com/google/uuid"
	"github.com/kouame-florent/axone-api/internal/axone"
	"gorm.io/gorm"
)

type AttachmentRepo struct {
	DB *gorm.DB
}

func NewAttachmentRepo(db *gorm.DB) *AttachmentRepo {
	return &AttachmentRepo{
		DB: db,
	}
}

func (r *AttachmentRepo) Create(e *axone.Attachment) (uuid.UUID, error) {
	tx := r.DB.Create(e)
	if tx.Error != nil {
		return uuid.UUID{}, tx.Error
	}
	return e.ID, nil
}

func (r *AttachmentRepo) Find(id uuid.UUID) (*axone.Attachment, error) {
	e := &axone.Attachment{}
	tx := r.DB.First(e, id)
	if tx.Error != nil {
		return &axone.Attachment{}, tx.Error
	}
	
	return e, nil
}

func (r *AttachmentRepo) FindAll() ([]axone.Attachment, error) {
	var ents []axone.Attachment
	tx := r.DB.Find(&ents)
	if tx.Error != nil {
		return []axone.Attachment{}, tx.Error
	}
	return ents, nil
}

func (r *AttachmentRepo) FindRange(offset, size int) ([]axone.Attachment, error) {
	return []axone.Attachment{}, nil
}

func (r *AttachmentRepo) Count() int {
	return 0
}

func (r *AttachmentRepo) Update(l *axone.Attachment) error {
	return nil
}

func (r *AttachmentRepo) Delete(l *axone.Attachment) error {
	return nil
}

