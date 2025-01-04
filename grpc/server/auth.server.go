package server

import (
	"context"
	authsvc "github.com/ngdangkietswe/swe-auth-service/grpc/service/auth"
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
	return s.authService.RegisterUser(ctx, req)
}

// Login is a function that implements the Login method of the AuthServiceServer interface
func (s *AuthGrpcServer) Login(ctx context.Context, req *auth.LoginReq) (*auth.LoginResp, error) {
	return s.authService.Login(ctx, req)
}

// EnableOrDisable2FA is a function that implements the EnableOrDisable2FA method of the AuthServiceServer interface
func (s *AuthGrpcServer) EnableOrDisable2FA(ctx context.Context, req *auth.EnableOrDisable2FAReq) (*auth.EnableOrDisable2FAResp, error) {
	return s.authService.EnableOrDisable2FA(ctx, req)
}

// ChangePassword is a function that implements the ChangePassword method of the AuthServiceServer interface
func (s *AuthGrpcServer) ChangePassword(ctx context.Context, req *auth.ChangePasswordReq) (*common.EmptyResp, error) {
	return s.authService.ChangePassword(ctx, req)
}

// ForgotPassword is a function that implements the ForgotPassword method of the AuthServiceServer interface
func (s *AuthGrpcServer) ForgotPassword(ctx context.Context, req *auth.ForgotPasswordReq) (*common.EmptyResp, error) {
	return s.authService.ForgotPassword(ctx, req)
}

// ResetPassword is a function that implements the ResetPassword method of the AuthServiceServer interface
func (s *AuthGrpcServer) ResetPassword(ctx context.Context, req *auth.ResetPasswordReq) (*common.EmptyResp, error) {
	return s.authService.ResetPassword(ctx, req)
}
