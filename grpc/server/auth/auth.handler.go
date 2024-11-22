package auth

import (
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
	"github.com/ngdangkietswe/swe-auth-service/data/repository/impl"
	authservice "github.com/ngdangkietswe/swe-auth-service/grpc/service/auth"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
	"google.golang.org/grpc"
)

type GrpcHandler struct {
	entClient *ent.Client
}

func NewGrpcHandler(entClient *ent.Client) *GrpcHandler {
	return &GrpcHandler{
		entClient: entClient,
	}
}

func (h *GrpcHandler) RegisterGrpcServer(server *grpc.Server) {
	authRepository := impl.NewAuthRepository(h.entClient)
	authService := authservice.NewAuthGrpcService(authRepository)
	auth.RegisterAuthServiceServer(server, NewGrpcServer(authService))
}
