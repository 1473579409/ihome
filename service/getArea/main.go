package main

import (
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"ihome-client/service/getArea/handler"
	"ihome-client/service/getArea/model"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	getArea "ihome-client/service/getArea/proto/getArea"
)

func main() {

	model.InitDb()
	model.InitRedis()
	// New Service
	consulRegistry := consul.NewRegistry()

	service := micro.NewService(
		micro.Name("go.micro.srv.getArea"),
		micro.Version("latest"),
		micro.Registry(consulRegistry),
	)

	// Initialise service
	service.Init()

	// Register Handler
	getArea.RegisterGetAreaHandler(service.Server(), new(handler.GetArea))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
