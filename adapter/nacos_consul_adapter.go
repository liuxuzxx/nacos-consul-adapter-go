package adapter

import (
	"log"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func NacosClient() {
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "172.16.1.15",
			Port:   8848,
		},
	}
	namingClient, err := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig": constant.ClientConfig{
			TimeoutMs:           5000,
			ListenInterval:      10000,
			NotLoadCacheAtStart: true,
			LogDir:              "data/nacos/log",
		},
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	instances, err := namingClient.SelectAllInstances(vo.SelectAllInstancesParam{
		GroupName:   "",
		ServiceName: "fgmp-service-base",
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("查看实例数量:%d\n", len(instances))
	for _, instance := range instances {
		log.Printf("服务的详情数据信息:%v\n", instance)
	}

	//获取Nacos上面所有的配置信息
	clientConfig := constant.ClientConfig{
		Endpoint:            "172.16.1.15:8848",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "data/nacos/log",
		CacheDir:            "cache/config",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}

	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	configPage, err := client.GetConfig(vo.ConfigParam{
		DataId: "fgmp-server-backend-dev.yml",
		Group:  "DEFAULT_GROUP",
	})
	if err != nil {
		log.Println("请求出现错误")
		log.Fatal(err.Error())
	}
	log.Printf("查看配置信息:%v\n", configPage)
}
