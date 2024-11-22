package auth

import (
	authservice "github.com/ngdangkietswe/swe-auth-service/grpc/service/auth"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
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
