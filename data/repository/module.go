package repository

import (
	"github.com/ngdangkietswe/swe-auth-service/data/datasource"
	actionrepo "github.com/ngdangkietswe/swe-auth-service/data/repository/action"
	authrepo "github.com/ngdangkietswe/swe-auth-service/data/repository/auth"
	permissionrepo "github.com/ngdangkietswe/swe-auth-service/data/repository/permission"
	resourcerepo "github.com/ngdangkietswe/swe-auth-service/data/repository/resource"
	userpermissionsrepo "github.com/ngdangkietswe/swe-auth-service/data/repository/user_permissions"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	datasource.NewEntClient,
	actionrepo.NewActionRepository,
	authrepo.NewAuthRepository,
	resourcerepo.NewResourceRepository,
	permissionrepo.NewPermissionRepository,
	userpermissionsrepo.NewUserPermissionsRepository,
)
