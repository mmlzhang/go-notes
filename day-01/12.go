package main

import "fmt"

// 声明一个函数类型
type cb func(int) int


func main(){
	testCallBack(1, callBack)
	testCallBack(2, func(x int) int {
		fmt.Println("回调函数，%d\n", x)
		return x
	})
}


func testCallBack(x int, f cb){
	fmt.Println("testCallBack, %d\n", x)
	f(x)
}


func callBack(x int) int {
	fmt.Println("callBack, x: %d\n", x)
	return x
}
