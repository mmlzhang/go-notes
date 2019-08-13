package main

import "sync"

// 顺序一致内存模型

//var a string
//var done bool
//
//func setup() {
//	a = "hello world"
//	done = true
//}
//
//func main() {
//	go setup()
//	for !done {}
//	print(a)
//}

//func main() {
//	done := make(chan int)
//	go func() {
//		println("hello world")
//		done <-1
//	}()
//
//	<-done
//}

func main() {
	var mu sync.Mutex

	mu.Lock()
	go func() {
		println("hhhh")
		mu.Unlock()
	}()
	mu.Lock()
}
