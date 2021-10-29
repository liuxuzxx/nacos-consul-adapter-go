package consul

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

var (
	adapter = NacosConsulAdapter{}
)

type NacosConsulAdapter struct {
	namingClient naming_client.INamingClient
}

func (n *NacosConsulAdapter) FetchNacosServices() Services {
	serviceList, err := n.loadNacosServices()
	if err != nil {
		log.Fatal(err.Error())
	}
	return ConvertServices(serviceList)
}

func (n *NacosConsulAdapter) loadNacosServices() (model.ServiceList, error) {
	serviceList, err := n.namingClient.GetAllServicesInfo(vo.GetAllServiceInfoParam{
		NameSpace: "public",
		GroupName: "DEFAULT_GROUP",
		PageNo:    1,
		PageSize:  20,
	})
	return serviceList, err
}

func (n *NacosConsulAdapter) FetchByServiceName(serviceName string) []Instance {
	sources, err := n.namingClient.SelectAllInstances(vo.SelectAllInstancesParam{
		GroupName:   "",
		ServiceName: serviceName,
	})
	if err != nil {
		log.Print(err.Error())
		return []Instance{}
	}
	return ConvertInstances(sources)
}

func (n *NacosConsulAdapter) FetchAgentInformation() string {
	return Agent
}

func (n *NacosConsulAdapter) HealthCheck(serviceName string) []Health {
	url := fmt.Sprintf("http://172.16.16.46:8500/v1/health/service/%s", serviceName)
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()

	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("查看从consul获取的数据信息:%s\n", string(responseBytes))
	healths := []Health{}
	json.Unmarshal(responseBytes, &healths)
	log.Printf("查看获取的数据解析:%v\n", healths)
	return healths
}

func InitNacosAdapter() NacosConsulAdapter {
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
	adapter.namingClient = namingClient
	return adapter
}
