package internal

import (
	"fmt"

	"github.com/RaymondCode/simple-demo/global"
	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
// Author [SliverHorn](https://github.com/SliverHorn)
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf 格式化打印日志
// Author [SliverHorn](https://github.com/SliverHorn)
func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch global.DY_CONFIG.System.DbType {
	case "mysql":
		logZap = global.DY_CONFIG.Mysql.LogZap
	case "pgsql":
		logZap = global.DY_CONFIG.Pgsql.LogZap
	}
	if logZap {
		global.DY_LOG.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
