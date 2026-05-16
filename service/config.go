package service

import (
	"bishe/model"
	"flag"
	"os"

	"gopkg.in/yaml.v2"
)

var Cfg *model.Config

// 加载配置文件
func LoadConfig() (err error) {
	configFile := flag.String("conf", "./config/config.yaml", "config文件")

	flag.Parse()

	Cfg = &model.Config{}
	data, err := os.ReadFile(*configFile)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(data, Cfg)
	return
}
