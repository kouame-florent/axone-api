package repo

/*
import (

	"github.com/google/uuid"
	"github.com/kouame-florent/axone-api/internal/axone"
	"gorm.io/gorm"
)

type OrganizationRepo struct {
	DB *gorm.DB
}

func NewOrganizationRepo(db *gorm.DB) *OrganizationRepo {
	return &OrganizationRepo{
		DB: db,
	}
}

func (r *OrganizationRepo) Create(e *axone.Organization) (uuid.UUID, error) {
	tx := r.DB.Create(e)
	if tx.Error != nil {
		return uuid.UUID{}, tx.Error
	}
	return e.ID, nil
}

func (r *OrganizationRepo) Find(id uuid.UUID) (*axone.Organization, error) {
	e := &axone.Organization{}
	tx := r.DB.First(e, id)
	if tx.Error != nil {
		return &axone.Organization{}, tx.Error
	}

	return e, nil
}

func (r *OrganizationRepo) FindAll() ([]axone.Organization, error) {
	var ents []axone.Organization
	tx := r.DB.Find(&ents)
	if tx.Error != nil {
		return []axone.Organization{}, tx.Error
	}
	return ents, nil
}

func (r *OrganizationRepo) FindRange(offset, size int) ([]axone.Organization, error) {
	return []axone.Organization{}, nil
}

func (r *OrganizationRepo) Count() int {
	return 0
}

func (r *OrganizationRepo) Update(l *axone.Organization) error {
	return nil
}

func (r *OrganizationRepo) Delete(l *axone.Organization) error {
	return nil
}

*/
