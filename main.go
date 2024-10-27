package main

import (
	"gvb-server/core"
	"gvb-server/global"
	"gvb-server/routers"
)

func main() {
	// 从配置文件读取配置
	core.InitConf()
	// 初始化日志
	global.LOG = core.InitLogger()
	// 连接数据库
	global.DB = core.InitGorm()
	// 初始化路由
	router := routers.InitRouter()
	addr := global.CONFIG.System.Addr()
	global.LOG.Infoln("gvb_server运行在: %s", addr)
	router.Run(addr)
}
