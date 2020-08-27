package main

import (
	"github.com/julienschmidt/httprouter"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/web"
	"net/http"
	"sss/IhomeWeb/handler"

	//"sss/IhomeWeb/handler"
	// micro/v2的注册器
	"github.com/micro/go-micro/v2/registry"
	// micro/v2的consul支持
	"github.com/micro/go-plugins/registry/consul/v2"
	//
	_"sss/IhomeWeb/models"
)

func main() {
	// 新建consul注册器
	consulReg := consul.NewRegistry(
		// 注册的consul信息
		registry.Addrs("127.0.0.1:8500"),
	)
	// create new web service
	service := web.NewService(
		web.Name("go.micro.web.IhomeWeb"),
		web.Version("latest"),
		// 服务添加consul支持
		web.Registry(consulReg),
		// web服务监听的ip和port
		web.Address(":8080"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}
	//使用路由中间件来映射页面
	rou := httprouter.New()
	rou.NotFound = http.FileServer(http.Dir("html"))
	//获取地区请求
	rou.GET("/api/v1.0/areas",handler.GetArea)
	//获取session
	rou.GET("/api/v1.0/session",handler.GetSession)
	//获取首页伦播图
	rou.GET("/api/v1.0/house/index",handler.GetIndex)
	//获取验证码图片
	rou.GET("/api/v1.0/imagecode/:uuid",handler.GetImageCd)
	//获取短信验证码
	rou.GET("/api/v1.0/smscode/:mobile",handler.GetSmscd)
	//登陆
	rou.POST("/api/v1.0/users",handler.PostRet)



	// register html handler
	service.Handle("/", rou)


	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
