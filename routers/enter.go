package routers

import (
	"github.com/gin-gonic/gin"
	"gvb-server/global"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.CONFIG.System.Env)
	router := gin.Default()
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	// 系统配置api
	routerGroupApp.SettingsRouter()
	return router
}
