package main

import (
	"fmt"
	"github.com/lvchengchang/canalysis/proxy"
	"github.com/spf13/cobra"
)

func main() {
	var caCmd = &cobra.Command{
		Use:   "analysis",
		Short: "canalysis",
		Long:  "canalysis",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("13")
		},
	}

	var cmdPrint = &cobra.Command{
		Use:   "run",
		Short: "canalysis run",
		Long:  `canalysis run`,
		Run: func(cmd *cobra.Command, args []string) {
			proxy.DeamonProxy(cmd, args)
		},
	}

	caCmd.AddCommand(cmdPrint)
	err := caCmd.Execute()
	if err != nil {
		return
	}
}
