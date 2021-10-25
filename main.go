package main

import (
	"log"
	"nacos_consul_adapter/adapter"
)

func main() {
	log.Println("初始化启动项目")
	adapter.NacosClient()
}
