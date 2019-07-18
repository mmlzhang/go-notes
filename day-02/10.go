package main

import "fmt"

/*
带缓冲的通道允许发送段的数据发送和接收段的数据处于异步状态，就是说发送端发送的数据可以放在缓冲去里面．
可以等待接收端去获取数据，而不是需要立刻需要接收端去获取数据
不过由于缓冲区的大小是有限的，所以还是必须有接收段来接收数据的，否则缓冲区一满，数据发送端就没法继续发送数据了

*/

// 遍历通道和关闭通道
func fib(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	ch <- 32
	fmt.Println(ch)

	c := make(chan int, 10)
	fmt.Println(cap(c), "cap(c)")
	go fib(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}

}
