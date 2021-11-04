package rest

import (
	"encoding/json"
	"log"
	"nacos_consul_adapter/consul"

	"github.com/kataras/iris/v12"
)

//放置伪装consul的提供http伪装请求的服务接口

var (
	Consul  = ConsulPretenderRest{}
	Adapter = consul.Adapter
)

type ConsulPretenderRest struct {
}

func (c *ConsulPretenderRest) FetchServiceByName(ctx iris.Context) {
	serviceName := ctx.Params().Get("serviceName")
	instances := Adapter.FetchByServiceName(serviceName)
	ctx.JSON(instances)
}

func (c *ConsulPretenderRest) FetchAllServices(ctx iris.Context) {
	services := Adapter.FetchNacosServices()
	ctx.JSON(services)
}

func (c *ConsulPretenderRest) FetchAgentInformation(ctx iris.Context) {
	agent := Adapter.FetchAgentInformation()
	ctx.WriteString(agent)
}

func (c *ConsulPretenderRest) FetchHealth(ctx iris.Context) {
	serviceName := ctx.Params().Get("serviceName")
	healths := Adapter.HealthCheck(serviceName)

	bytes, _ := json.Marshal(healths)
	log.Printf("查看伪装结果：%s\n", string(bytes))
	ctx.Header("X-Consul-Index", "10796")
	ctx.JSON(healths)
}
