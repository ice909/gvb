package routers

import (
	"github.com/gin-gonic/gin"
	"gvb-server/global"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.CONFIG.System.Env)
	router := gin.Default()
	router.GET("", func(context *gin.Context) {
		context.JSON(200, gin.H{})
	})
	return router
}
