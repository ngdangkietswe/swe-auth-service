package auth

import (
	"context"
	authservice "github.com/ngdangkietswe/swe-auth-service/grpc/service/auth"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/common"
)

type GrpcServer struct {
	auth.UnimplementedAuthServiceServer
	authService authservice.IAuthService
}

func NewGrpcServer(authService authservice.IAuthService) *GrpcServer {
	return &GrpcServer{
		authService: authService,
	}
}

// RegisterUser is a function that implements the RegisterUser method of the AuthServiceServer interface
func (s *GrpcServer) RegisterUser(ctx context.Context, req *auth.User) (*common.UpsertResp, error) {
	return s.authService.RegisterUser(ctx, req)
}

// Login is a function that implements the Login method of the AuthServiceServer interface
func (s *GrpcServer) Login(ctx context.Context, req *auth.LoginReq) (*auth.LoginResp, error) {
	return s.authService.Login(ctx, req)
}
