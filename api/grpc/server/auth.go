package server

import (
	"encoding/base64"
	"log"
	"strings"

	"github.com/kouame-florent/axone-api/internal/axone"
	"github.com/kouame-florent/axone-api/internal/repo"
	"github.com/kouame-florent/axone-api/internal/svc"
)

func (s *AxoneServer) valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Basic ")
	userRep := repo.NewUserRepo(s.DB)
	userSvc := svc.NewUserSvc(userRep)

	decToken, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return false
	}
	log.Printf("Decoded token: %s", decToken)
	creds := strings.Split(string(decToken), ":")
	if len(creds) != 2 {
		return false
	}

	login := creds[0]
	passwd := creds[1]

	err = userSvc.Valid(axone.Credential{Login: login, Password: passwd}, axone.USER_STATUS_ENABLED)
	return err == nil
}

func (s *AxoneServer) authenticate(creds axone.Credential) (axone.UserProfile, error) {

	userRep := repo.NewUserRepo(s.DB)
	userSvc := svc.NewUserSvc(userRep)

	profile, err := userSvc.Authenticate(creds, axone.USER_STATUS_ENABLED)
	if err != nil {
		return axone.UserProfile{}, err
	}

	return profile, nil
}
