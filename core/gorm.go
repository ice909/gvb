package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvb-server/global"
	"log"
	"time"
)

func InitGorm() *gorm.DB {
	if global.CONFIG.Mysql.Host == "" {
		log.Println("未配置mysql,取消gorm连接")
		return nil
	}
	dsn := global.CONFIG.Mysql.Dsn()
	var mysqlLogger logger.Interface
	if global.CONFIG.System.Env == "dev" {
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: mysqlLogger})
	if err != nil {
		log.Fatalf(fmt.Sprintf("[%s] mysql连接失败", dsn))
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour * 4)
	return db
}
