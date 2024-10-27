package api

import "gvb-server/api/settings"

type ApiGroup struct {
	Settings settings.SettingsApi
}

var ApiGroupApp = new(ApiGroup)
