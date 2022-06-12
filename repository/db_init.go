package repository

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	var err error
<<<<<<< HEAD
<<<<<<< HEAD
	fmt.Println("hhh")
	dsn := "root:123456@tcp(127.0.0.1:3306)/dy_base?charset=utf8mb4&parseTime=True&loc=Local"
=======
	dsn := "root:root@tcp(127.0.0.1:3306)/dy_database?charset=utf8mb4&parseTime=True&loc=Local"
>>>>>>> a5ad9421cddcb4c71a3ebda7d6ed77f835c4b828
=======

	dsn := "root:root@tcp(127.0.0.1:3306)/dy_database?charset=utf8mb4&parseTime=True&loc=Local"

>>>>>>> upstream/gzh
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}
