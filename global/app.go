package global

import (
	"github.com/RaymondCode/simple-demo/config"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/sessions"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	DY_ConfigViper   *viper.Viper  `json:"config_viper,omitempty"`
	DY_CONFIG        config.Server `json:"config"`
	DY_LOG           *zap.Logger   `json:"dy___log,omitempty"`
	DY_DB            *gorm.DB
	DY_DBList        map[string]*gorm.DB
	DY_JWTMW         *jwt.GinJWTMiddleware
	DY_REDIS         *redis.Client
	DY_SESSION_STORE sessions.Store
}

var App = new(Application)
