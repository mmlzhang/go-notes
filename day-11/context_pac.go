package main

import (
	"sync"
	"context"
	"fmt"
	"time"
)

/*
context: 用来简化对于处理单个请求的多个Goroutine之间与请求域的数据，超时和退出等操作
*/

// 使用context包来重新实现前面的线程安全退出或超时的控制
func worker(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	for {
		select {
		default:
			fmt.Println("hello")
		    time.Sleep(time.Second)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}


func main() {
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//
	//var wg sync.WaitGroup
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go worker(ctx, &wg)
	//}
	//time.Sleep(5*time.Second)
	//cancel()
	//wg.Wait()

	// 通过 Context　控制后台Goroutine状态
	ctx, cancel := context.WithCancel(context.Background())

	ch := GenerateNature(ctx)
	for i:=0; i<100;i++ {
		prime := <-ch
		fmt.Printf("%v: %#v", i+1, prime)
		ch = PrimeFilter(ctx, ch, prime)
	}
	cancel()
}

/*
go 语言是自带内存自动回收的特性的，因此内存一般不会泄露，在前面素数筛的列子中，
GenerateNatural和PermeFilter函数内部都启动了新的Goroutine,当main函数不在使用管道时后台
Goroutinue有泄露的风险，我们可以使用context包来避免这个问题
*/
// 返回生成自然数序列的管道：　２，　３，　４
func GenerateNature(ctx, context.Context) chan int {
	ch := make(chan int)
	go func() {
		for i :=2;;i++ {
			select {
			case <- ctx.Done():
				return

			case ch<-i:


			}
		}
	}()
	return ch
}

// 管道过滤器: 删除能被素数整除的数
func PrimeFilter(ctx context.Context, in <- chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		if i := <-in;i%prime !=0{
			select{
			case <- ctx.Done():
				return
			case out <- i:

			}
		}
	}()
	return out
}