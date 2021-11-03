package main

import (
	"log"
	"nacos_consul_adapter/rest"
	"nacos_consul_adapter/started/cmd"
	"os"
	"time"

	"github.com/kataras/iris/v12"
)

func main() {
	cmd.Execute()
	app := iris.New()
	f, _ := os.Create("iris.log")
	app.Logger().SetOutput(f)
	app.Logger().SetLevelOutput("error", f)
	app.Use(Logger())
	consulAPI := app.Party("/v1")
	{
		consulAPI.Use(iris.Compression)
		consulAPI.Get("/catalog/service/{serviceName}", rest.Consul.FetchServiceByName)
		consulAPI.Get("/catalog/services", rest.Consul.FetchAllServices)
		consulAPI.Get("/agent/self", rest.Consul.FetchAgentInformation)
		consulAPI.Get("/health/service/{serviceName}", rest.Consul.FetchHealth)
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
