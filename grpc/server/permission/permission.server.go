package permission

import (
	permissionsvc "github.com/ngdangkietswe/swe-auth-service/grpc/service/permission"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
)

type GrpcServer struct {
	auth.UnimplementedPermissionServiceServer
	permissionSvc permissionsvc.IPermissionService
}

func NewGrpcServer(permissionSvc permissionsvc.IPermissionService) *GrpcServer {
	return &GrpcServer{
		permissionSvc: permissionSvc,
	}
}
