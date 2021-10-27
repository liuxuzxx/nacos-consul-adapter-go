package main

import (
	"log"
	"nacos_consul_adapter/adapter"
	"nacos_consul_adapter/consul"

	"github.com/kataras/iris/v12"
)

func main() {
	log.Println("初始化启动项目")
	adapter.NacosClient()

	//开始引入Iris来做模拟consul的操作
	app := iris.New()
	consulAPI := app.Party("/v1")
	{
		consulAPI.Use(iris.Compression)
		consulAPI.Get("/catalog/service/{serviceName}", consul.Consul.FetchServiceByName)
	}
	app.Listen(":18500")
}
