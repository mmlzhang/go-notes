package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct {

}

func (nokia NokiaPhone) call() {
	fmt.Println("I am Nokia, I ca call you!")
}

type IPhone struct {

}

func (iphone IPhone) call() {
	fmt.Println("I am iphone, I can call you!")
}


func main() {
	/*　接口处理　*/
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
