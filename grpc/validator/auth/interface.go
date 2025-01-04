package auth

import (
	"context"
	"github.com/google/wire"
	authrepo "github.com/ngdangkietswe/swe-auth-service/data/repository/auth"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
)

type IAuthValidator interface {
	RegisterUser(ctx context.Context, req *auth.User) error
	ChangePassword(req *auth.ChangePasswordReq, hashCurrentPassword string) error
}

// ProvideAuthValidator is a function to provide an auth validator
func ProvideAuthValidator() (a IAuthValidator) {
	wire.Build(
		authrepo.ProvideAuthRepository(),
		NewAuthValidator,
	)
	return
}
