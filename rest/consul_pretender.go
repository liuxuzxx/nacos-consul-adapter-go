package rest

import "github.com/kataras/iris/v12"

//放置伪装consul的提供http伪装请求的服务接口

var (
	Consul = ConsulPretenderRest{}
)

type ConsulPretenderRest struct {
}

func (c *ConsulPretenderRest) FetchServiceByName(ctx iris.Context) {
	serviceName := ctx.Params().Get("serviceName")
	serviceDetails := map[string]string{
		serviceName: "Mastering Concurrency in Go",
	}
	ctx.JSON(serviceDetails)
}
