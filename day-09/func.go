package main

import "fmt"

// func Add(a, b int) int {
// 	return a+b
// }

// // 匿名函数
// var Add = func(a, b, int) int {
// 	return a + b
// }


// // 多个参数和多个返回值
// func Swap(a, b int) (int, int) {
// 	return a, b
// }


// // 可变数量的参数
// func Sum(a int, more ...int) int {
// 	for _, v := range more {
// 		a += v 
// 	}
// 	return a
// }


func Print(a ...interface{}) {
	fmt.Println(a...)
}


// // 返回值也可以进行命名
// // 如果返回值命名了，可以通过名字修改返回志，也可以通过defer语句在reture语句之后修改返回值
// func Find(m, map[int]int, key int) (value int, ok bool) {
// 	value, ok = m[key]
// 	return
// }

// fun Inc() (v int) {
// 	defer func() { v++ }()
// 	return 42
// }

// 当可变参数是一个空接口类型时，调用者是否解包可变参数会导致不同结果
func main() {
	// var a = []interface{}{123, "abc"}
	// Print(a...)
	// Print(a)

	// tmp := Inc()
	// fmt.Println("tmp value: ", tmp)

	// for i := 0; i < 3; i++ {
	// 	defer func() {fmt.Println(i)}()
	// }

	// for i := 0; i<3; i++ {
	// 	i := i
	// 	defer func() { println(i) } ()
	// }

	for i := 0; i < 3; i ++ {
		fmt.Println("i = ", i)
		defer func(i int) { println(i)}(i)
	}
}
