package initialize

import (
	"context"
	"github.com/RaymondCode/simple-demo/global"
	"github.com/go-redis/redis/v8"
)

func InitializeRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     global.App.Config.Redis.Addr,
		Password: global.App.Config.Redis.Password, // no password set
		DB:       global.App.Config.Redis.DB,       // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil
	}
	return client
}
