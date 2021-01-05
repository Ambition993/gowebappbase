package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var rdb *redis.Client
var ctx = context.Background()

type ReidsConfig struct {
	Host     string `mapstructure:"mysql.host"`
	Port     int
	Password string
	DB       int
	PoolSize int
}

func Init() (err error) {
	var redisConfig = ReidsConfig{
		Host:     viper.GetString("redis.host"),
		Port:     viper.GetInt("redis.port"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.pool_size"),
	}
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			redisConfig.Host,
			redisConfig.Port,
		),
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
		PoolSize: redisConfig.PoolSize,
	})
	_, err = rdb.Ping(ctx).Result()
	return err
}
func Close() {
	_ = rdb.Close()
}
