package main

import (
	"fmt"
	"go_learn/day-10/pubsub"
	"time"
	"strings"
)

func Producer(factor int, out chan<-int) {
	for i := 0; ; i++ {
		out <- i*factor
	}
}


func Consumer(in <-chan int) {
	for v := range in{
		fmt.Println(v)
	}
}


func main() {
	//ch := make(chan int, 64)
	//
	//go Producer(3, ch)
	//go Producer(5, ch)
	//go Consumer(ch)
	//
	////time.Sleep(time.Second * 1)
	//
	//// Ctrl + C 退出
	//sig := make(chan os.Signal, 1)
	//signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	//fmt.Printf("quit  (%v)\n", <-sig)


	p := pubsub.NewPublisher(100*time.Millisecond, 10)
	defer p.Close()


	all := p.Subscribe()
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok :=  v.(string); ok {
			return strings.COntains(s, "golang")
		}
		return false
	})

	p.Publish("hello world")
	p.Publish("hello world")

	go func() {
		for msg := range all {
			fmt.Println("all", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println("golang", msg)
		}
	}()


	time.Sleep(time.Second * 3)

}
