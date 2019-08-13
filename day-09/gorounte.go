package main

//var done = make(chan bool)
//var msg string
//
//func aGoroutine() {
//	msg = "hello world"
//	done <- true
//}
//
//func main() {
//	go aGoroutine()
//	<- done
//	println(msg)
//}

// 使用close(c) 关闭管道代替 done <- false

//var done = make(chan bool)
//var msg string
//
//func aGoroutine() {
//	msg = "hello world 1"
//	close(done)
//
//}
//
//func main() {
//	go aGoroutine()
//	<- done
//	println(msg)
//}


var limit = make(chan int, 3)

func main() {
	for _, w := range work {
		go func() {
			limit <- 1
			w()
			<- limit
		}()

	}
	select {}
}

// 使用同步的思路: 使用显示的同步