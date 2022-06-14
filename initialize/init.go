package initialize

import (
	_ "database/sql"
	"errors"
	"github.com/RaymondCode/simple-demo/global"
)

func Init() error {
	// init Viper
	InitializeConfig()
	// init Redis
	global.App.DY_REDIS = InitializeRedis()
	// init zap log
	global.App.DY_LOG = InitializeLog()
	// init gorm and connect db
	global.App.DY_DB = Gorm()
	if global.App.DY_DB == nil {
		return errors.New("gorm initialize failed")
	}
	// init tables
	RegisterTables(global.App.DY_DB)

	return nil
}
