package service

import (
	grpcserviceauth "github.com/ngdangkietswe/swe-auth-service/grpc/service/auth"
	grpcservicepermission "github.com/ngdangkietswe/swe-auth-service/grpc/service/permission"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		grpcserviceauth.NewAuthGrpcService,
		grpcservicepermission.NewPermissionGrpcService,
	),
)
