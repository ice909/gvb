package main

import (
	"gvb-server/core"
	"gvb-server/global"
)

func main() {
	// 从配置文件读取配置
	core.InitConf()
	// 初始化日志
	global.LOG = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()
}
