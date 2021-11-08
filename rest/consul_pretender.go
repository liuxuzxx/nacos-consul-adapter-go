package rest

import (
	"encoding/json"
	"nacos_consul_adapter/consul"

	log "github.com/sirupsen/logrus"

	"github.com/kataras/iris/v12"
)

//放置伪装consul的提供http伪装请求的服务接口

type ConsulPretenderRest struct {
	Adapter consul.NacosConsulAdapter
}

func (c *ConsulPretenderRest) FetchServiceByName(ctx iris.Context) {
	serviceName := ctx.Params().Get("serviceName")
	instances := c.Adapter.FetchByServiceName(serviceName)
	ctx.Header("X-Consul-Index", "10796")
	ctx.JSON(instances)
}

func (c *ConsulPretenderRest) FetchAllServices(ctx iris.Context) {
	services := c.Adapter.FetchNacosServices()
	ctx.Header("X-Consul-Index", "10796")
	ctx.JSON(services)
}

func (c *ConsulPretenderRest) FetchAgentInformation(ctx iris.Context) {
	agent := c.Adapter.FetchAgentInformation()
	ctx.Header("X-Consul-Index", "10796")
	ctx.WriteString(agent)
}

func (c *ConsulPretenderRest) FetchHealth(ctx iris.Context) {
	serviceName := ctx.Params().Get("serviceName")
	healths := c.Adapter.HealthCheck(serviceName)

	bytes, _ := json.Marshal(healths)
	log.Printf("查看伪装结果：%s\n", string(bytes))
	ctx.Header("X-Consul-Index", "10796")
	ctx.JSON(healths)
}
