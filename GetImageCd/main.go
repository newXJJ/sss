package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	//"github.com/micro/micro/v2/cmd/protoc-gen-micro/plugin/micro"
	"sss/GetImageCd/handler"

	GetImageCd "sss/GetImageCd/proto/GetImageCd"
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
		micro.Name("go.micro.service.GetImageCd"),
		micro.Version("latest"),
		// 服务添加consul支持
		micro.Registry(consulReg),
	)

	// Initialise service
	service.Init()

	// Register Handler
	GetImageCd.RegisterGetImageCdHandler(service.Server(), new(handler.GetImageCd))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.service.GetImageCd", service.Server(), new(subscriber.GetImageCd))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
