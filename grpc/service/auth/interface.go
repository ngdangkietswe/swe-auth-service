package auth

import (
	"context"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/common"
)

type IAuthService interface {
	RegisterUser(ctx context.Context, req *auth.User) (*common.UpsertResp, error)
	Login(ctx context.Context, req *auth.LoginReq) (*auth.LoginResp, error)
	EnableOrDisable2FA(ctx context.Context, req *auth.EnableOrDisable2FAReq) (*auth.EnableOrDisable2FAResp, error)
	ChangePassword(ctx context.Context, req *auth.ChangePasswordReq) (*common.EmptyResp, error)
}
