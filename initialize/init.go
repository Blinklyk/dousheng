package initialize

import (
	"errors"
	"github.com/RaymondCode/simple-demo/global"
)

func Init() error {
	// init Viper
	InitializeConfig()
	// init Redis
	global.DY_REDIS = InitializeRedis()
	// init zap log
	//zap.ReplaceGlobals(global.DY_LOG)
	// init gorm and connect db
	global.DY_DB = Gorm()
	if global.DY_DB == nil {
		return errors.New("gorm initialize failed")
	}
	// init tables
	RegisterTables(global.DY_DB)

	return nil
}
