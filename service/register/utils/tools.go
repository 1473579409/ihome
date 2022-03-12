package utils

import (
	"github.com/asim/go-micro/plugins/registry/consul/v4"

	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
)

func GetMicroClient() client.Client {
	consulReg := consul.NewRegistry()
	microService := micro.NewService(
		micro.Registry(consulReg),
	)
	return microService.Client()
}
