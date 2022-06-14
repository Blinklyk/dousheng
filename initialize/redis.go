package initialize

import (
	"context"
	"github.com/RaymondCode/simple-demo/global"
	"github.com/go-redis/redis/v8"
)

func InitializeRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     global.App.DY_CONFIG.Redis.Addr,
		Password: global.App.DY_CONFIG.Redis.Password, // no password set
		DB:       global.App.DY_CONFIG.Redis.DB,       // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.App.DY_LOG.Error(err.Error())
		return nil
	}
	return client
}
