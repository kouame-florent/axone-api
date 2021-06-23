package svc

import (
	"github.com/kouame-florent/axone-api/internal/axone"
	"github.com/kouame-florent/axone-api/internal/repo"
	"gorm.io/gorm"
)

type UserSvc struct {
	Repo *repo.UserRepo
}

func NewUserSvc(r *repo.UserRepo) *UserSvc {
	return &UserSvc{
		Repo: r,
	}
}

func (u *UserSvc) Authenticate(cred axone.Credential, status axone.UserStatus) (bool, error) {
	tx := u.Repo.DB.Where("login = ? AND password = ? AND status = ?", cred.Login, cred.Password, status)
	if tx.Error == gorm.ErrRecordNotFound {
		return false, nil
	}

	if tx.Error != nil {
		return false, tx.Error
	}
	return true, nil
}
