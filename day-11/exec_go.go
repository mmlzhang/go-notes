package main

import (
	"fmt"
	"time"
)

func main() {
	// 基于select实现管道的超时判断
	select {
	case v := <-in:
		fmt.Println(v)
	case <-time.After(time.Second):
		return // 超时
	}

	// 通过select的default分支实现非阻塞的管道发送或接收操作
	select {
	case v := in:
		fmt.Println(v)
	default:

	}

	// 通过select来组织main函数退出
	//func main () {
	//	select {
	//
	//	}
	//}
}