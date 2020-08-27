package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"sss/GetArea/handler"

	GetArea "sss/GetArea/proto/GetArea"
	// micro/v2的注册器
	"github.com/micro/go-micro/v2/registry"
	// micro/v2的consul支持
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	// 新建consul注册器
	consulReg := consul.NewRegistry(
		// 注册的consul信息
		registry.Addrs("127.0.0.1:8500"),
	)

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.GetArea"),
		micro.Version("latest"),
		// 服务添加consul支持
		micro.Registry(consulReg),
	)

	// Initialise service
	service.Init()

	// Register Handler
	GetArea.RegisterGetAreaHandler(service.Server(), new(handler.GetArea))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.service.GetArea", service.Server(), new(subscriber.GetArea))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
