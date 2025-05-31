package server

import (
	"context"
	authsvc "github.com/ngdangkietswe/swe-auth-service/grpc/service/auth"
	"github.com/ngdangkietswe/swe-go-common-shared/util"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/common"
)

type AuthGrpcServer struct {
	auth.UnimplementedAuthServiceServer
	authService authsvc.IAuthService
}

func NewAuthGrpcServer(authService authsvc.IAuthService) *AuthGrpcServer {
	return &AuthGrpcServer{
		authService: authService,
	}
}

// RegisterUser is a function that implements the RegisterUser method of the AuthServiceServer interface
func (s *AuthGrpcServer) RegisterUser(ctx context.Context, req *auth.User) (*common.UpsertResp, error) {
	return util.HandleGrpc(ctx, req, s.authService.RegisterUser)
}

// Login is a function that implements the Login method of the AuthServiceServer interface
func (s *AuthGrpcServer) Login(ctx context.Context, req *auth.LoginReq) (*auth.LoginResp, error) {
	return util.HandleGrpc(ctx, req, s.authService.Login)
}

// EnableOrDisable2FA is a function that implements the EnableOrDisable2FA method of the AuthServiceServer interface
func (s *AuthGrpcServer) EnableOrDisable2FA(ctx context.Context, req *auth.EnableOrDisable2FAReq) (*auth.EnableOrDisable2FAResp, error) {
	return util.HandleGrpc(ctx, req, s.authService.EnableOrDisable2FA)
}

// ChangePassword is a function that implements the ChangePassword method of the AuthServiceServer interface
func (s *AuthGrpcServer) ChangePassword(ctx context.Context, req *auth.ChangePasswordReq) (*common.EmptyResp, error) {
	return util.HandleGrpc(ctx, req, s.authService.ChangePassword)
}

// ForgotPassword is a function that implements the ForgotPassword method of the AuthServiceServer interface
func (s *AuthGrpcServer) ForgotPassword(ctx context.Context, req *auth.ForgotPasswordReq) (*common.EmptyResp, error) {
	return util.HandleGrpc(ctx, req, s.authService.ForgotPassword)
}

// ResetPassword is a function that implements the ResetPassword method of the AuthServiceServer interface
func (s *AuthGrpcServer) ResetPassword(ctx context.Context, req *auth.ResetPasswordReq) (*common.EmptyResp, error) {
	return util.HandleGrpc(ctx, req, s.authService.ResetPassword)
}
