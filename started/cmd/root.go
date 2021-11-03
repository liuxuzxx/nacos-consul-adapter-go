package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nacos-consul-adapter-go config ./config.yml",
	Short: "nacos-consul-adapter-go是一个适配Nacos为Consul数据格式的适配器，专门为Prometheus提供Consul伪装服务接口转换器.",
	Long:  `由于Prometheus和Promtail只有四种服务发现：手动，文件服务发现，Consul服务发现，K8S服务发现，并没有Nacos的服务发现配置，但是系统内使用的又是Nacos，所以需要适配器来伪装Nacos为Consul!`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal("错误，请输入config和及其配置文件路径的参数")
		os.Exit(-1)
	},
}

func Execute() {
	rootCmd.Execute()
}
