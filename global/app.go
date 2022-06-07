package global

import (
	"github.com/RaymondCode/simple-demo/config"
	"github.com/spf13/viper"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Server
	Redis       config.Redis
}

var App = new(Application)
