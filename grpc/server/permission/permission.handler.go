package permission

import (
	"github.com/ngdangkietswe/swe-auth-service/grpc/service/permission"
	grpcauth "github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
	"google.golang.org/grpc"
)

type GrpcHandler struct {
}

func NewGrpcHandler() *GrpcHandler {
	return &GrpcHandler{}
}

func (h *GrpcHandler) RegisterGrpcServer(server *grpc.Server) {
	permissionSvc := permission.ProvidePermissionService()

	grpcauth.RegisterPermissionServiceServer(server, NewGrpcServer(permissionSvc))
	grpcauth.RegisterPermissionInternalServiceServer(server, NewGrpcInternalServer(permissionSvc))
}
