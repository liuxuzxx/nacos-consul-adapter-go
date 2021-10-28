package rest

import (
	"nacos_consul_adapter/consul"

	"github.com/kataras/iris/v12"
)

//放置伪装consul的提供http伪装请求的服务接口

var (
	Consul = ConsulPretenderRest{}
)

type ConsulPretenderRest struct {
}

func (c *ConsulPretenderRest) FetchServiceByName(ctx iris.Context) {
	serviceName := ctx.Params().Get("serviceName")
	instances := consul.Adapter.FetchByServiceName(serviceName)
	ctx.JSON(instances)
}

func (c *ConsulPretenderRest) FetchAllServices(ctx iris.Context) {
	consul.Adapter.FetchNacosServices()
	ctx.JSON("ok，获取成功")
}
