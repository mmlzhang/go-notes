package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	var (
		i int
		wg sync.WaitGroup
	)
	wg.Add(2)
	go func() {
		for {
			i++
			time.Sleep(time.Nanosecond)  // sleep ++生效
		}
	}()
	go func() {
		for {
			fmt.Println(i)
		}
	}()
	wg.Wait()
}