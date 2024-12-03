package main

import (
	"github.com/ngdangkietswe/swe-auth-service/grpc"
	"github.com/ngdangkietswe/swe-go-common-shared/config"
)

func main() {
	config.Init()
	grpc.NewGrpcServer()
}
