package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config.path",
	Short: "帮助信息.",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("请输入一个配置文件的绝对地址参数")
			os.Exit(-1)
		}
		fmt.Printf("查看参数信息%s 长度:%d\n", args[0], len(args))
		//加载配置文件
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
