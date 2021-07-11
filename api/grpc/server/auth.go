package server

import (
	"context"
	"encoding/base64"
	"strings"

	"github.com/kouame-florent/axone-api/api/grpc/gen"
	"github.com/kouame-florent/axone-api/internal/axone"
	"github.com/kouame-florent/axone-api/internal/repo"
	"github.com/kouame-florent/axone-api/internal/svc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

//ensureValidBasicCredentials is a grpc interceptor to validate basic credentials
func (s *AxoneServer) ensureValidBasicCredentials(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}
	if !s.valid(md["authorization"]) {
		return nil, errInvalidToken
	}
	return handler(ctx, req)
}

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
	//log.Printf("Decoded token: %s", decToken)
	creds := strings.Split(string(decToken), ":")
	if len(creds) != 2 {
		return false
	}

	login := creds[0]
	passwd := creds[1]

	err = userSvc.Validate(axone.Credential{Login: login, Password: passwd}, axone.USER_STATUS_ENABLED)
	return err == nil
}

//Login authenticate end users and return profile information
func (s *AxoneServer) Login(ctx context.Context, req *gen.LoginRequest) (*gen.LoginResponse, error) {

	profile, err := s.userInfos(axone.Credential{Login: req.Login, Password: req.Password})
	if err != nil {
		return &gen.LoginResponse{}, err
	}

	lr := &gen.LoginResponse{
		UserID:    profile.UserID,
		Login:     profile.Login,
		Password:  profile.Password,
		Email:     profile.Email,
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
	}

	return lr, err
}

func (s *AxoneServer) userInfos(creds axone.Credential) (axone.UserProfile, error) {

	userRep := repo.NewUserRepo(s.DB)
	userSvc := svc.NewUserSvc(userRep)

	profile, err := userSvc.GetUserInfos(creds, axone.USER_STATUS_ENABLED)
	if err != nil {
		return axone.UserProfile{}, err
	}

	return profile, nil
}
