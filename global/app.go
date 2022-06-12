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
	DY_ConfigViper   *viper.Viper          `json:"dy_config_viper,omitempty"`
	DY_CONFIG        config.Server         `json:"dy_config"`
	DY_LOG           *zap.Logger           `json:"dy_log,omitempty"`
	DY_DB            *gorm.DB              `json:"dy_db"`
	DY_DBList        map[string]*gorm.DB   `json:"dy_db_list"`
	DY_JWTMW         *jwt.GinJWTMiddleware `json:"dy_jwtmw"`
	DY_REDIS         *redis.Client         `json:"dy_redis"`
	DY_SESSION_STORE sessions.Store        `json:"dy_session_store"`
}

var App = new(Application)
