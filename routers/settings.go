package routers

import (
	"gvb-server/api"
)

func (router *RouterGroup) SettingsRouter() {
	router.GET("settings", api.ApiGroupApp.Settings.SettingsInfoView)
}
