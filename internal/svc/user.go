package svc

import (
	"fmt"

	"github.com/google/uuid"
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

// GetUserInfos retrieve user information from  database
func (u *UserSvc) GetUserInfos(cred axone.Credential, status axone.UserStatus) (axone.UserProfile, error) {
	var profile axone.UserProfile
	tx := u.Repo.DB.Model(&axone.User{}).Select("users.id as user_id,users.login,users.password,users.email,users.first_name,users.last_name").
		Where("login = ? AND password = ? AND status = ? ", cred.Login, cred.Password, status).Scan(&profile)
	if tx.Error == gorm.ErrRecordNotFound {
		return axone.UserProfile{}, fmt.Errorf("%s", "authentication failed")
	}

	if tx.Error != nil {
		return axone.UserProfile{}, tx.Error
	}
	return profile, nil
}

//check if user exist and has given status in database
func (u *UserSvc) Validate(cred axone.Credential, status axone.UserStatus) error {
	//	var profil axone.UserProfil
	tx := u.Repo.DB.Model(&axone.User{}).Select("1").
		Where("login = ? AND password = ? AND status = ? ", cred.Login, cred.Password, status)
	if tx.Error == gorm.ErrRecordNotFound {
		return fmt.Errorf("%s", "authentication failed")
	}

	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (u *UserSvc) CreateUser(user *axone.User) (uuid.UUID, error) {
	return u.Repo.Create(user)

}
