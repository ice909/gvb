package core

import (
	"gopkg.in/yaml.v3"
	"gvb-server/global"
	"log"
	"os"
)

// InitConf 读取yaml文件配置
func InitConf() {
	const ConfigFile = "settings.yaml"
	c := &global.CONFIG
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config Init success")
}
