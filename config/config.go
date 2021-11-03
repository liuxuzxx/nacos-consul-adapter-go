package config

import (
	"encoding/json"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	NacosConfigs    []NacosConfig
	NameSpaceGroups []NameSpaceGroup
}

type NacosConfig struct {
	IP   string
	Port uint64
}

type NameSpaceGroup struct {
	NameSpace  string
	GroupNames []string
}

var Conf Config

func InitConfig(configPath string) {
	log.Println("开始加载配置信息....")
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yml")
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("更新配置文件")
	})
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("解析配置文件出错!")
	}
	viper.Unmarshal(&Conf)
	jsonString, _ := json.Marshal(&Conf)
	log.Printf("查看加载的config配置信息:%s\n", string(jsonString))
}
