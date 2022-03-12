package main

import (
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"ihome-client/service/userOrder/handler"

	"github.com/asim/go-micro/plugins/registry/consul/v4"
	userOrder "ihome-client/service/userOrder/proto/userOrder"

	"ihome-client/service/userOrder/model"
)

func main() {
	// New Service
	consulReg := consul.NewRegistry()

	service := micro.NewService(
		micro.Name("go.micro.srv.userOrder"),
		micro.Version("latest"),
		micro.Registry(consulReg),
		micro.Address(":9986"),
	)

	// Initialise service
	service.Init()
	model.InitDb()

	// Register Handler
	userOrder.RegisterUserOrderHandler(service.Server(), new(handler.UserOrder))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
