package auth

import (
	"context"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/common"
)

type IAuthService interface {
	RegisterUser(ctx context.Context, req *auth.User) (*common.UpsertResp, error)
	Login(ctx context.Context, req *auth.LoginReq) (*auth.LoginResp, error)
}
