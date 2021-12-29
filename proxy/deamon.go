package proxy

import (
	"github.com/lvchengchang/canalysis/conf"
	"github.com/lvchengchang/canalysis/log"
	"github.com/lvchengchang/canalysis/mysql"
	"github.com/spf13/cobra"
)

func DeamonProxy(cmd *cobra.Command, args []string) {
	conf.InitConf()
	log.InitLogger()

	mysql.NewConnection().Connect()
}
