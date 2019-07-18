// go语言常量

package main

import "fmt"

func main()  {
	const LENGTH int = 10
	const WIDTH int =5
	var area int
	const a, b, c = 1, false, "str"  // 多重赋值
	area = LENGTH * WIDTH
	fmt.Println("面积为： %d", area)
	println()
	println(a, b, c)
}

// 常量还可以做枚举
const (
	Unkwnown = 0
	Femail = 1
	Male = 2
)