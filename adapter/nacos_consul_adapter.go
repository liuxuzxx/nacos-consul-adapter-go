package adapter

import (
	"log"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func NacosClient() {
	client, err := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": []constant.ServerConfig{
			{
				IpAddr: "172.16.1.15",
				Port:   8848,
			},
		},
		"clientConfig": constant.ClientConfig{
			TimeoutMs:           5000,
			ListenInterval:      10000,
			NotLoadCacheAtStart: true,
			LogDir:              "data/nacos/log",
			NamespaceId:         "a4495738-12c0-42b3-a036-82d3002bdd7a",
		},
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	serviceList, err := client.GetAllServicesInfo(vo.GetAllServiceInfoParam{
		NameSpace: "public",
		GroupName: "DEFAULT_GROUP",
		PageNo:    1,
		PageSize:  20,
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("获取的服务数量是:%d\n", serviceList.Count)
	for _, item := range serviceList.Doms {
		log.Printf("服务:%s\n", item)
		instances, err := client.SelectAllInstances(vo.SelectAllInstancesParam{
			ServiceName: item,
		})
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Printf("查看实例数量:%d\n", len(instances))
		for _, temp := range instances {
			log.Printf("查看具体详情:%v\n", temp)
		}
	}
}
