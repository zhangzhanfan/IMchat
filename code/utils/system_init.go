package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB = InitConfigMYSQL()
	// RD = InitConfigREDIS()
	Red *redis.Client
)

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

//初始化redis
func InitConfigREDIS() {
	ctx := context.Background()
	Red = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.pollSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
	result, err := Red.Ping(ctx).Result()
	if err != nil {
		fmt.Println("init redis fail......", err)
	} else {
		fmt.Println("init redis success......", result)
	}
}

const (
	PublishKey = "websocket"
)

// publish 发布消息redis
func Publish(ctx context.Context, channel string, msg string) error {
	// var err error
	err := Red.Publish(ctx, channel, msg).Err()
	fmt.Println("publish a msg ", msg)
	return err
}

//订阅消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	fmt.Println("写了吗")
	sub := Red.Subscribe(ctx, channel)
	for {
		msg, err := sub.Receive(ctx)
		switch msg.(type) {
		case *redis.Subscription:
			fmt.Println("订阅成功")
		case *redis.Message:
			m := msg.(redis.Message)
			fmt.Println("接受到一条消息 ", m.Payload)
			return m.Payload, err
		case *redis.Pong:
			fmt.Println("收到pong了")
		default:
			fmt.Println("报错了...default")
		}
		if err != nil {
			fmt.Println("报错了...", err)
		}
	}
}
