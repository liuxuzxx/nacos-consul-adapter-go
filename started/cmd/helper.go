package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helperCmd = &cobra.Command{
	Use:   "help",
	Short: "帮助信息.",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("帮助信息")
	},
}

func init() {
	rootCmd.AddCommand(helperCmd)
}
