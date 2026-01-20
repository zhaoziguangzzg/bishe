package service

import (
	"bishe/internal/app/knowledge_sharing/conf"
	"flag"
	"os"

	"gopkg.in/yaml.v2"
)

var Cfg *conf.Config

// 加载配置文件
func LoadConfig() (err error) {
	configFile := flag.String("conf", "./configs/config.yaml", "config文件")

	flag.Parse()

	Cfg = &conf.Config{}
	data, err := os.ReadFile(*configFile)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(data, Cfg)
	return
}
