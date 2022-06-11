package initialize

import (
	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

func Gorm() *gorm.DB {
	m := global.App.DY_CONFIG.Mysql
	// TODO use switch and case as below comment shows
	db, err := gorm.Open(mysql.New(mysql.Config{
		//DSN:                       "root:root@tcp(127.0.0.1:3306)/dy_database?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name
		DSN:                       m.Dsn(),
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		SkipDefaultTransaction: false, // 默认false
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "dy_", // 添加表名前缀
			SingularTable: true,  // 启用单数表名，
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 是否关闭逻辑外键 （代码自动外键关系）
	})
	if err != nil {
		global.App.DY_LOG.Error(err.Error())
		return nil
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(m.MaxIdleConns) // 设置最大连接数
	sqlDB.SetMaxOpenConns(m.MaxOpenConns) // 设置最大可空闲连接数
	return db

}

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		// 测试用户表
		model.User{},
		model.Video{},
		model.Favorite{},
		model.Comment{},
		model.Follow{},
		model.Follower{},
	)
	if err != nil {
		global.App.DY_LOG.Error(err.Error())
		os.Exit(0)
	}
	global.App.DY_LOG.Info("register table success")
}
