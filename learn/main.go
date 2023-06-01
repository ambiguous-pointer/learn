package main

import (
	"fmt"
	"github.com/ambiguous-pointer/learn-go/two/pojo"
	"github.com/ambiguous-pointer/learn-go/two/utils"
)

func main() {
	utils.SelectUsers()
	users := []pojo.SysUsers{
		{Password: "password", Username: "ad"},
		{Password: "password1", Username: "ad1"},
	}
	for index, value := range users {
		fmt.Println(index)
		fmt.Println(value)
	}
	println(utils.InsertUser(users))
}
