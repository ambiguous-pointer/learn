package collback

import (
	"fmt"
)

// Phone 定义接口
type Phone interface {
	call()
	call2(phone Phone) string
}

type NokiaPhone struct {
	id           int
	name         string
	categoryId   int
	categoryName string
}

func (nokiaPhone NokiaPhone) call2(phone Phone) string {
	return "this is nokiaPhone collBack"
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
	memberId       int
	memberBalance  float32
	memberSex      bool
	memberNickname string
}

func (iPhone IPhone) call2(phone Phone) string {
	return "this is iPhone collBack"
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func Test01() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
