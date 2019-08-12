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
	/*
	第一种防范是 在循环体内部在定义一个局部变量,这样每次迭代 defer 语句的闭包函数捕获的都是不同的变量,
	这些变量的只对应迭代的值,第二种方法是讲地带变量通过闭包函数的参数传入,defer语句会马上对调用参数求值,两种方式都是可以工作的,不过一般来说,在for循环内部执行defer语句并不是一个好习惯
	go语言中,如果以切片为参数调用函数时,有时候会给人一种参数采用了引用的方式的假象,因为在被调用函数的内部可以修改传入的切片元素,
	其实,任何可以通过函数参数修改调用参数的情形,都是应为函数参数中显示或隐式传入了指针参数,函数参数传值的规范
	准确说是只针对数据结构中固定部分传值.
	*/

	x := []int{1000, 100, 222}
	twice(x)
	fmt.Println(x)

	y := IntSliceHeader{Data: []int{11, 22, 33}, Len:3, Cap:111}
	twice2(y)

	fmt.Println(y)
	/*
	因为切片中的底层数组部分是通过隐式指针传递
	*/

}

func twice(x []int) {
	for i := range x {
		x[i] *= 2
	}
}

type IntSliceHeader struct {
	Data []int
	Len int
	Cap int
}

func twice2(x IntSliceHeader) {
	for i := 0; i < x.Len; i ++ {
		x.Data[i] *= 2
	}
}
