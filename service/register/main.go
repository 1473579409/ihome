package main

import (
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"ihome-client/service/register/handler"

	"github.com/asim/go-micro/plugins/registry/consul/v4"
	register "ihome-client/service/register/proto/register"

	"ihome-client/service/register/model"
)

func main() {
	//服务发现用consul
	consulReg := consul.NewRegistry()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.register"),
		micro.Version("latest"),
		micro.Registry(consulReg),
		micro.Address(":9982"),
	)

	// Initialise service
	service.Init()
	model.InitRedis()
	model.InitDb()

	// Register Handler
	register.RegisterRegisterHandler(service.Server(), new(handler.Register))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
