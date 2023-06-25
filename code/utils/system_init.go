package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB = InitConfigMYSQL()

// var Log = InitLogrus() //日志对象

//初始化配置
func InitConfig() {
	viper.AddConfigPath("/config")
	viper.SetConfigName("app")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("初始化配置#################失败！", err)
	}
}

//初始化MySQL
func InitConfigMYSQL() *gorm.DB {
	// InitConfig()
	//自定义日志模板，打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	database, err := gorm.Open(mysql.Open(viper.GetString("mysql.dsn")), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}
	// println("连接数据库成功")
	return database
}
