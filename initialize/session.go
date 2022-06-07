package initialize

import (
	"github.com/RaymondCode/simple-demo/global"
	"github.com/gin-contrib/sessions/redis"
	"log"
)

func InitSession() {
	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	if err != nil {
		panic("session init error")
	}
	global.DY_SESSION_STORE = store
	log.Println("init session successfully")
}
