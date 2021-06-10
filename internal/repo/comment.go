
package repo

import (

	"github.com/google/uuid"
	"github.com/kouame-florent/axone-api/internal/axone"
	"gorm.io/gorm"
)

type CommentRepo struct {
	DB *gorm.DB
}

func NewCommentRepo(db *gorm.DB) *CommentRepo {
	return &CommentRepo{
		DB: db,
	}
}

func (r *CommentRepo) Create(e *axone.Comment) (uuid.UUID, error) {
	tx := r.DB.Create(e)
	if tx.Error != nil {
		return uuid.UUID{}, tx.Error
	}
	return e.ID, nil
}

func (r *CommentRepo) Find(id uuid.UUID) (*axone.Comment, error) {
	e := &axone.Comment{}
	tx := r.DB.First(e, id)
	if tx.Error != nil {
		return &axone.Comment{}, tx.Error
	}
	
	return e, nil
}

func (r *CommentRepo) FindAll() ([]axone.Comment, error) {
	var ents []axone.Comment
	tx := r.DB.Find(&ents)
	if tx.Error != nil {
		return []axone.Comment{}, tx.Error
	}
	return ents, nil
}

func (r *CommentRepo) FindRange(offset, size int) ([]axone.Comment, error) {
	return []axone.Comment{}, nil
}

func (r *CommentRepo) Count() int {
	return 0
}

func (r *CommentRepo) Update(l *axone.Comment) error {
	return nil
}

func (r *CommentRepo) Delete(l *axone.Comment) error {
	return nil
}

