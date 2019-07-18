// 指针　　＆　＊
package main

import "fmt"

func main(){
	var a int = 2
	var b int32
	var c float32
	var ptr *int

	/* 运算符实例 */
	fmt.Printf("第一行，％T\n", a)
	fmt.Printf("第2行，％T\n", b)
	fmt.Printf("第3行，％T\n", c)

	fmt.Printf("第4行，％T\n", ptr)

	ptr = &a
	fmt.Printf("a值：　%d\n", a)
	fmt.Printf("*ptr 为　%d\n", *ptr)
	
}