package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"proxy_pool/app/global"
	"proxy_pool/app/models"
	"proxy_pool/config"
)

var DB *gorm.DB

func init()  {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", config.CONFIG.Mysql.Username, config.CONFIG.Mysql.Password, config.CONFIG.Mysql.Host,
	config.CONFIG.Mysql.Port, config.CONFIG.Mysql.Dbname)


	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		global.Logger.Errorln("Can't not connect database.")
	}
	DB.AutoMigrate(&models.ProxyModel{})
}