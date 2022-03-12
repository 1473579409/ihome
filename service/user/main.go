package main

import (
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"ihome-client/service/user/handler"

	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"ihome-client/service/user/model"
	user "ihome-client/service/user/proto/user"
)

func main() {
	//使用consul做服务发现
	consulReg := consul.NewRegistry()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
		micro.Registry(consulReg),
		micro.Address(":9984"),
	)

	// Initialise service
	service.Init()
	model.InitDb()

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
