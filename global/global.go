package global

import (
	"gorm.io/gorm"
	"gvb-server/config"
)

var (
	CONFIG *config.Config
	DB     *gorm.DB
)
