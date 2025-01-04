package server

import (
	"context"
	permissionsvc "github.com/ngdangkietswe/swe-auth-service/grpc/service/permission"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/common"
)

type PermissionInternalGrpcServer struct {
	auth.UnimplementedPermissionInternalServiceServer
	permissionSvc permissionsvc.IPermissionService
}

func NewPermissionInternalGrpcServer(permissionSvc permissionsvc.IPermissionService) *PermissionInternalGrpcServer {
	return &PermissionInternalGrpcServer{
		permissionSvc: permissionSvc,
	}
}

func (s *PermissionInternalGrpcServer) PermissionOfUser(ctx context.Context, req *common.IdReq) (*auth.PermissionOfUserResp, error) {
	return s.permissionSvc.PermissionOfUser(ctx, req)
}
