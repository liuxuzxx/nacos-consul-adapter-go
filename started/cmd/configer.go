package cmd

import (
	"fmt"
	"log"
	"nacos_consul_adapter/config"
	"nacos_consul_adapter/consul"
	"os"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "帮助信息.",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("请输入一个配置文件的绝对地址参数")
			os.Exit(-1)
		}
		fmt.Printf("查看参数信息%s 长度:%d\n", args[0], len(args))
		config.InitConfig(args[0])
		consul.InitNacosAdapter(config.Conf)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
