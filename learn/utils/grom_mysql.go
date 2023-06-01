package utils

import (
	"fmt"
	"github.com/ambiguous-pointer/learn-go/two/mysql"
	"github.com/ambiguous-pointer/learn-go/two/pojo"
)

func SelectUsers() {
	db, err := mysql.GetDB()
	if err != nil {
		return
	}
	var users []pojo.SysUsers
	result := db.Find(&users).Table("sys_users")
	affected := result.RowsAffected
	fmt.Println(affected)
	for index, value := range users {
		fmt.Println(index)
		fmt.Println(value)
	}
	err = result.Error // returns error
	if err != nil {
		return
	} else {
		return
	}
}
func InsertUser(users []pojo.SysUsers) int64 {
	db, err := mysql.GetDB()
	if err != nil {
		return 0
	}
	db.Create(&users)
	return db.RowsAffected
}
