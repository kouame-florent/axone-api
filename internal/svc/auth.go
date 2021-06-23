package svc

import (
	"encoding/base64"
	"fmt"
	"log"
	"strings"
)

func Valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Basic ")
	return token == base64.StdEncoding.EncodeToString([]byte("homer:homer"))
}

func checkDB(token string) error {
	decToken, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return err
	}
	log.Printf("Decoded token: %s", decToken)
	creds := strings.Split(string(decToken), ":")
	if len(creds) != 2 {
		return fmt.Errorf("%s", "cannot decode auth token")
	}

	//	login := creds[0]
	//	passwd := creds[1]

	return nil
}
