package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Nacos-Consul-Adapter的版本号",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deveops-databack-1.0.0版本")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
