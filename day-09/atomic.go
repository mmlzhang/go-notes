package main

import (
	"sync"
	"sync/atomic"
	"fmt"
	"time"
)

var total2 uint64

func worker2(wg *sync.WaitGroup) {
	defer wg.Done()
	var i uint64
	for i = 0; i <= 100; i++ {
		atomic.AddUint64(&total2, i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker2(&wg)
	go worker2(&wg)
	wg.Wait()

	fmt.Println(total2)
}


type singleton struct {}

var (
	instance *singleton
	initialized  uint32
	mu      sync.Mutex

)


func Instance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = &singleton{}
	}
	return instance
}


type Once struct {
	m Mutex
	done uint32
}


func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}

	o.m.Lock()
	defer o.m.Unlock()

	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}


// 基于sync.Once实现单件模式
var (
	instance *singleton
	once sync.Once

)

func Instance() *singleton {
	once.Do(func() {
		instance = &singleton
	})
	return instance
}


var config atomic.Value

config.Store(loadConfig())

go func() {
	for {
		time.Sleep(time.Second)
		config.Store(loadConfig)
	}
}()

for i := 0; i < 10; i++ {
	go func() {
		for r := range requests(){
			c := config.Load()
}
}
}