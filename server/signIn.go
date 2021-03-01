package server

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/zhashkevych/auth/pkg/auth/usecase"
	"sport_kld/app/models/user_model"
)

type UserParams struct {
	Login string
	Password string
}

func SignIn(ctx context.Context, userParams UserParams) (string, error) {

	pwd := sha1.New()
	pwd.Write([]byte(userParams.Password))
	pwd.Write([]byte("saaaaalt"))

	password := fmt.Sprintf("%x", pwd.Sum(nil))


}
