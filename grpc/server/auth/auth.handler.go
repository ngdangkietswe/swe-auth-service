package auth

import (
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
	"github.com/ngdangkietswe/swe-auth-service/data/repository/impl"
	service "github.com/ngdangkietswe/swe-auth-service/grpc/service/auth"
	validator "github.com/ngdangkietswe/swe-auth-service/grpc/validator/auth"
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
	authValidator := validator.NewAuthValidator(authRepository)
	authService := service.NewAuthGrpcService(authRepository, authValidator)
	auth.RegisterAuthServiceServer(server, NewGrpcServer(authService))
}
