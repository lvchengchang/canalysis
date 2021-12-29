package conf

import (
	"fmt"
	"github.com/lvchengchang/canalysis/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Mysql
}

type Mysql struct {
	Addr     string
	Port     string
	Username string
	Password string
}

var Conf *Config

var conflog = log.Logger.With()

func InitConf() {
	config, err := ioutil.ReadFile("./canalysis.yaml")
	if err != nil {
		conflog.Fatal(fmt.Sprintf("init config err: %s", err.Error()))
	}
	err = yaml.Unmarshal(config, &Conf)
	if err != nil {
		conflog.Fatal(fmt.Sprintf("unmarshal config err: %s", err.Error()))
	}
}
