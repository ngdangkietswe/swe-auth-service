package auth

import (
	"context"
	"github.com/google/wire"
	authrepo "github.com/ngdangkietswe/swe-auth-service/data/repository/auth"
	validator "github.com/ngdangkietswe/swe-auth-service/grpc/validator/auth"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/common"
)

type IAuthService interface {
	RegisterUser(ctx context.Context, req *auth.User) (*common.UpsertResp, error)
	Login(ctx context.Context, req *auth.LoginReq) (*auth.LoginResp, error)
	EnableOrDisable2FA(ctx context.Context, req *auth.EnableOrDisable2FAReq) (*auth.EnableOrDisable2FAResp, error)
	ChangePassword(ctx context.Context, req *auth.ChangePasswordReq) (*common.EmptyResp, error)
}

// ProvideAuthService is a function to provide an auth service
func ProvideAuthService() (a IAuthService) {
	wire.Build(
		authrepo.ProvideAuthRepository(),
		validator.ProvideAuthValidator,
		NewAuthGrpcService,
	)
	return
}
