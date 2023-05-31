package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB() (db *gorm.DB, err error) {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	db, err = gorm.Open(mysql.Open("root:password@tcp(127.0.0.1:3306)/gin-admin?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据出现错误")
	}
	return db, nil
}
