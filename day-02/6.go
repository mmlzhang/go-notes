package main

import "fmt"

func Factorial(n uint64) (result uint64) {
	if (n > 0) {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	// 递归函数 阶乘
	var i int = 15
	fmt.Printf("%d　的阶乘是　%d\n", i, Factorial(uint64(i)))

	for i := 0; i < 10; i++ {
		fmt.Printf("%d\t", fib(i))
	}

	// 类型转换
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("mean : %f\n", mean)
}
