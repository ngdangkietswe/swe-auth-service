package permission

import (
	"context"
	permissionsvc "github.com/ngdangkietswe/swe-auth-service/grpc/service/permission"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/common"
)

type GrpcInternalServer struct {
	auth.UnimplementedPermissionInternalServiceServer
	permissionSvc permissionsvc.IPermissionService
}

func NewGrpcInternalServer(permissionSvc permissionsvc.IPermissionService) *GrpcInternalServer {
	return &GrpcInternalServer{
		permissionSvc: permissionSvc,
	}
}

func (s *GrpcServer) PermissionOfUser(ctx context.Context, req *common.IdReq) (*auth.PermissionOfUserResp, error) {
	return s.permissionSvc.PermissionOfUser(ctx, req)
}
