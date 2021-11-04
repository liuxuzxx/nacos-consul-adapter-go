package main

import (
	"log"
	"nacos_consul_adapter/config"
	"nacos_consul_adapter/consul"
	"nacos_consul_adapter/rest"
	"nacos_consul_adapter/started/cmd"
	"os"
	"time"

	"github.com/kataras/iris/v12"
)

func main() {
	cmd.Execute()

	consulRest := rest.ConsulPretenderRest{
		Adapter: consul.InitNacosAdapter(config.Conf),
	}

	log.Println("开始初始化Iris-web服务")
	app := iris.New()
	f, _ := os.Create("iris.log")
	app.Logger().SetOutput(f)
	app.Logger().SetLevelOutput("error", f)
	app.Use(Logger())
	consulAPI := app.Party("/v1")
	{
		consulAPI.Use(iris.Compression)
		consulAPI.Get("/catalog/service/{serviceName}", consulRest.FetchServiceByName)
		consulAPI.Get("/catalog/services", consulRest.FetchAllServices)
		consulAPI.Get("/agent/self", consulRest.FetchAgentInformation)
		consulAPI.Get("/health/service/{serviceName}", consulRest.FetchHealth)
	}
	app.Listen(":18500")
}

func Logger() iris.Handler {
	return func(ctx iris.Context) {
		t := time.Now()
		ctx.Next()
		latency := time.Since(t)
		log.Print(latency)
		status := ctx.GetStatusCode()
		log.Println(status)
	}
}
