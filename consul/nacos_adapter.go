package consul

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	Adapter = NacosConsulAdapter{}
)

type NacosConsulAdapter struct {
}

func (n *NacosConsulAdapter) FetchNacosServices() {
	url := "http://172.16.16.46:8500/v1/catalog/services"
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
	services := Services{}
	json.Unmarshal(responseBytes, &services)
	log.Printf("查看获取的数据解析:%v\n", services)
}

func (n *NacosConsulAdapter) FetchByServiceName(serviceName string) []Instance {
	url := fmt.Sprintf("http://172.16.16.46:8500/v1/catalog/service/%s", serviceName)
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
	instances := []Instance{}
	log.Printf("查看获取的原始结果:%s\n", string(responseBytes))
	json.Unmarshal(responseBytes, &instances)
	log.Printf("查看获取的数据解析:%v\n", instances)
	return instances
}
