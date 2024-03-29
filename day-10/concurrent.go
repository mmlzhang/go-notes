package main

import "fmt"

/*
CSP  Communication Sequential Process  通讯顺序进程
有精确的数学模型

并发不是并行，并发跟跟更关注的是程序设计层面．并发的程序完全是可以顺序执行的，只有在真正的多核CPU上才能真正的同时运行
	并行更关注的是程序的运行层面，并行一般是简单的大量重复，例如CPU中对图像处理会有大量的并行运算．


go 并发方案
　不要通过共享内存来通信，⽽应通过通信来共享内存
*/

func main() {
	done := make(chan int, 10)

	for i := 0; i < cap(done); i++ {
		go func(i int) {
			fmt.Println("hello, world", i)
			done <- 1
		}(i)
	}

	for i := 0; i < cap(done); i++ {
		<- done
	}

}
