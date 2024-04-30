package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DB  *gorm.DB
	Rdb *redis.Client
)

// InitConfig 初始化viper
func InitConfig(fileName string) (err error) {
	viper.SetConfigFile(fileName)
	err = viper.ReadInConfig() //读取配置信息
	if err != nil {
		//读取配置信息失败
		fmt.Println("viper.ReadInConfig err:", err)
		return
	}
	fmt.Println("config inited 。。。")
	return
}

// InitMySQL 初始化MySQL
func InitMySQL() {
	//自定义日志模板 打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			LogLevel:      logger.Info, //级别
			Colorful:      true,        //彩色
		},
	)

	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dsn")),
		&gorm.Config{Logger: newLogger})
	fmt.Println(" MySQL inited 。。。")
	/*	user := models.UserBasic{}
		DB.Where("name = ?", "ls").First(&user)
		fmt.Println(user)*/
}

// InitRedis 初始化Redis
func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.db"),
		PoolSize:     viper.GetInt("redis.pool_size"),
		MinIdleConns: viper.GetInt("redis.min_idle_conns"),
	})
	_, err := Rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("init Redis err:", err)
	} else {
		fmt.Println("Redis inited 。。。")
	}
	/*	s, err := Rdb.Get("b").Result()
		fmt.Println(s)*/
}

const (
	PublishKey = "websocket" //发布频道
)

// Publish 发布消息到Redis的channel频道
func Publish(ctx context.Context, channel string, msg string) error {
	var err error
	fmt.Println("Publish 。。。。", msg)
	err = Rdb.Publish(ctx, channel, msg).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// Subscribe 订阅Redis频道并接收消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	//订阅频道
	sub := Rdb.Subscribe(ctx, channel)
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("Subscribe 。。。。", msg.Payload)
	return msg.Payload, err
}
