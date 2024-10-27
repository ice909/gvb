package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gvb-server/config"
)

var (
	CONFIG *config.Config
	DB     *gorm.DB
	LOG    *logrus.Logger
)
