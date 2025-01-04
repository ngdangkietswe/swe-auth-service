package service

import (
	grpcserviceauth "github.com/ngdangkietswe/swe-auth-service/grpc/service/auth"
	grpcservicepermission "github.com/ngdangkietswe/swe-auth-service/grpc/service/permission"
	"github.com/ngdangkietswe/swe-go-common-shared/cache"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		cache.NewRedisCache,
		grpcserviceauth.NewAuthGrpcService,
		grpcservicepermission.NewPermissionGrpcService,
	),
)
