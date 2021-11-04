package consul

import (
	"log"
	"nacos_consul_adapter/config"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

var (
	Adapter = NacosConsulAdapter{}
)

type NacosConsulAdapter struct {
	namingClient naming_client.INamingClient
	config       config.Config
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
	sources, err := n.namingClient.SelectAllInstances(vo.SelectAllInstancesParam{
		GroupName:   "",
		ServiceName: serviceName,
	})
	if err != nil {
		log.Print(err.Error())
		return []Health{}
	}
	return ConvertHealths(sources)
}

func (n *NacosConsulAdapter) initAdapter(config config.Config) {
	n.config = config
	serverConfigs := []constant.ServerConfig{}
	for _, value := range config.NacosConfigs {
		serverConfigs = append(serverConfigs, constant.ServerConfig{
			IpAddr: value.IP,
			Port:   value.Port,
		})
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
	n.namingClient = namingClient
}

func InitNacosAdapter(config config.Config) NacosConsulAdapter {
	Adapter.initAdapter(config)
	return Adapter
}
