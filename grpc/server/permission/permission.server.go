package permission

import (
	"context"
	permissionsvc "github.com/ngdangkietswe/swe-auth-service/grpc/service/permission"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/auth"
	"github.com/ngdangkietswe/swe-protobuf-shared/generated/common"
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

// UpsertPermission is a function that implements the UpsertPermission method of the PermissionServiceServer interface
func (s *GrpcServer) UpsertPermission(ctx context.Context, req *auth.UpsertPermissionReq) (*common.UpsertResp, error) {
	return s.permissionSvc.UpsertPermission(ctx, req)
}

// ListPermissions is a function that implements the ListPermissions method of the PermissionServiceServer interface
func (s *GrpcServer) ListPermissions(ctx context.Context, req *auth.ListPermissionsReq) (*auth.ListPermissionsResp, error) {
	return s.permissionSvc.ListPermissions(ctx, req)
}

// AssignPermissions is a function that implements the AssignPermissions method of the PermissionServiceServer interface
func (s *GrpcServer) AssignPermissions(ctx context.Context, req *auth.AssignPermissionsReq) (*common.EmptyResp, error) {
	return s.permissionSvc.AssignPermissions(ctx, req)
}
