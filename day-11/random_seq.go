package main

import (
	"fmt"
	"time"
	"sync"
)

func main () {
	//ch := make(chan int)
	//
	//go func() {
	//	for {
	//		select {
	//		case ch <- 0:
	//		case ch <- 1:
	//
	//		}
	//	}
	//}()
	//
	//for v := range ch {
	//	fmt.Println(v)
	//
	//}
	//channel1 := make(chan bool)
	//go worker(channel1)
	//time.Sleep(time.Second * 3)
	//channel1 <- false

	//channel2 := make(chan bool)
	//
	//for i := 0; i < 10; i ++ {
	//	go worker2(channel2)
	//}
	//time.Sleep(time.Second * 5)
	//close(channel2)

	channel3 := make(chan bool)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker3(&wg, channel3)
	}

	time.Sleep(time.Second * 10)
	close(channel3)
	wg.Wait()
}


// 通过select 和default分支实现一个Gorountine的退出控制
func worker(channel chan bool) {
	for {
		select {
		default:
			fmt.Println("hello")
			time.Sleep(time.Second)
		case <- channel:
			// 退出
		}
	}
}

/*
通过管道的发送操作来实现停止多个Goroutine需要创建同样数量的管道，代价太大
可以通过close关闭一个管道来实现广播效果，所有从关闭管道接收的操作都会收到一个零值和一个可选失败的标志位
*/

func worker2(channel chan bool) {
	for  {
		select {
		default:
			fmt.Println("hello success")
		    time.Sleep(time.Second)
		case <-channel:
			//退出
		}
	}
}


/*
通过close关闭channel执行Goroutine的广播退出功能，这个程序依然不够健壮，
当每个Goroutine收到退出指令退出时一般会进行一定的清理工作，但是退出的清理工作并不能被保证，因为
main线程并没有等待哥哥工作的Goroutine退出工作完成，可以结合sync.WaitGroup改善
*/

func worker3(wg *sync.WaitGroup, channel chan bool) {
	defer wg.Done()

	for {
		select {
		case <-channel:
			return
		default:
			fmt.Println("hello 3 ")
		    time.Sleep(time.Second)
		}
	}
}
