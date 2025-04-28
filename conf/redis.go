package conf

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"time"
)

// 局部变量用于后续封装进自定义结构体
var rDB *redis.Client

// 键值过期时间
var rDuration = 30 * 24 * time.Hour

type RedisClient struct {
}

func InitRedis() (*RedisClient, error) {
	rDB = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.dsn"),
		Password: "",
		DB:       0,
	})
	// 查看链接是否成功
	_, err := rDB.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return &RedisClient{}, err
}

// 简单封装

func (r *RedisClient) Set(key, value string) error {
	// 设置过期时间并返回错误
	return rDB.Set(context.Background(), key, value, rDuration).Err()
}

func (r *RedisClient) Get(key string) (any, error) {
	return rDB.Get(context.Background(), key).Result()
}

func (r *RedisClient) Delete(keys ...string) error {
	return rDB.Del(context.Background(), keys...).Err()
}
