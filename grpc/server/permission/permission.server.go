package permission

import (
	"context"
	permissionsvc "github.com/ngdangkietswe/swe-auth-service/grpc/service/permission"
	"github.com/ngdangkietswe/swe-go-common-shared/constants"
	"github.com/ngdangkietswe/swe-go-common-shared/domain"
	"github.com/ngdangkietswe/swe-go-common-shared/security"
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
	var action string
	if req.Id != nil {
		action = constants.ActionUpdate
	} else {
		action = constants.ActionCreate
	}
	return security.SecuredAuth(ctx, req, domain.Permission{Action: action, Resource: constants.ResourcePermission}, s.permissionSvc.UpsertPermission)
}

// ListPermissions is a function that implements the ListPermissions method of the PermissionServiceServer interface
func (s *GrpcServer) ListPermissions(ctx context.Context, req *auth.ListPermissionsReq) (*auth.ListPermissionsResp, error) {
	return security.SecuredAuth(ctx, req, domain.Permission{Action: constants.ActionRead, Resource: constants.ResourcePermission}, s.permissionSvc.ListPermissions)
}

// AssignPermissions is a function that implements the AssignPermissions method of the PermissionServiceServer interface
func (s *GrpcServer) AssignPermissions(ctx context.Context, req *auth.AssignPermissionsReq) (*common.EmptyResp, error) {
	return s.permissionSvc.AssignPermissions(ctx, req)
}
