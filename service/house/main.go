package main

import (
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"ihome-client/service/house/handler"

	"github.com/asim/go-micro/plugins/registry/consul/v4"
	house "ihome-client/service/house/proto/house"

	"ihome-client/service/house/model"
)

func main() {
	// New Service
	consulReg := consul.NewRegistry()

	service := micro.NewService(
		micro.Name("go.micro.srv.house"),
		micro.Version("latest"),
		micro.Address(":9985"),
		micro.Registry(consulReg),
	)

	// Initialise service
	service.Init()
	model.InitDb()

	// Register Handler
	house.RegisterHouseHandler(service.Server(), new(handler.House))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
