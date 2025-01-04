package main

import (
	"github.com/ngdangkietswe/swe-auth-service/data/repository"
	"github.com/ngdangkietswe/swe-auth-service/grpc"
	"github.com/ngdangkietswe/swe-go-common-shared/config"
	"go.uber.org/fx"
	grpcserver "google.golang.org/grpc"
)

func main() {
	config.Init()
	app := fx.New(
		repository.Module,
		grpc.Module,
		fx.Invoke(func(*grpcserver.Server) {}),
	)
	app.Run()
}
