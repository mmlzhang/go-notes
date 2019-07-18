package main

import (
	"time"
	"fmt"
)

// 并发

func say(s string){
	for i :=0; i< 2; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

/*
通道　channel 是用来传递数据的一个数据结构
通道可用在两个goroutine之间通过传递一个指定类型的值来同步运行和通讯，操作符为　<-
发送或者接收，如果未指定方向，则为双向通道
*/

func sum(s []int, c chan int){
	sum := 0
	for _, v := range s{
		sum += v
	}
	c <- sum  // 把sum发送到通道c
}

func main() {
	//go say("111111")
	say("22222222")

	s := []int{1, 2, 45, -1, -9, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

}
