package auth

import (
	"context"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
)

type IAuthValidator interface {
	RegisterUser(ctx context.Context, req *auth.User) error
	ChangePassword(req *auth.ChangePasswordReq, hashCurrentPassword string) error
	ResetPassword(req *auth.ResetPasswordReq) error
}
