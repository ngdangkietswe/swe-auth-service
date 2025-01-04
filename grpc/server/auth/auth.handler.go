package auth

import (
	"github.com/ngdangkietswe/swe-auth-service/grpc/service/auth"
	grpcauth "github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
	"google.golang.org/grpc"
)

type GrpcHandler struct {
}

func NewGrpcHandler() *GrpcHandler {
	return &GrpcHandler{}
}

func (h *GrpcHandler) RegisterGrpcServer(server *grpc.Server) {
	authService := auth.ProvideAuthService()
	grpcauth.RegisterAuthServiceServer(server, NewGrpcServer(authService))
}
