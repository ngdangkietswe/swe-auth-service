package permission

import (
	"github.com/ngdangkietswe/swe-auth-service/data/ent"
	"github.com/ngdangkietswe/swe-auth-service/data/repository/impl"
	service "github.com/ngdangkietswe/swe-auth-service/grpc/service/permission"
	validator "github.com/ngdangkietswe/swe-auth-service/grpc/validator/permission"
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
	authRepo := impl.NewAuthRepository(h.entClient)
	actionRepo := impl.NewActionRepository(h.entClient)
	resourceRepo := impl.NewResourceRepository(h.entClient)
	permissionRepo := impl.NewPermissionRepository(h.entClient)
	userPermissionsRepo := impl.NewUserPermissionsRepository(h.entClient)

	permissionValidator := validator.NewPermissionValidator(actionRepo, resourceRepo, permissionRepo, authRepo)

	permissionSvc := service.NewPermissionGrpcService(actionRepo, resourceRepo, permissionRepo, userPermissionsRepo, authRepo, permissionValidator)

	auth.RegisterPermissionServiceServer(server, NewGrpcServer(permissionSvc))
}