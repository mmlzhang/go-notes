package main

import (
	"fmt"
	"image"
	"io"
	"image/png"
	"image/jpeg"
)

func main() {
	fmt.Println("day-08")
	//var a [3]int
	//var b = [...]int{1, 2, 3}
	//var c = [...]int{2: 3, 1: 2}
	//var d = [...]int{1, 2, 4:5, 6}
	//fmt.Println(a, b, c, d)
	// 空位自动填充0

	//var a = [...]int{1, 2, 3}
	//var b = &a
	//fmt.Println(a[0], a[1])
	//fmt.Println(b[0], b[1]) // 通过数组指针访问数组元素和直接访问元素类似
	//// 可以通过for range 来迭代数组指针致指向的元素，　其实数组指针复制时会拷贝一个指针
	//// 但是数组指针不够用灵活，因为数组的长度时数组类型的组成部分，指向不同长度的数组指针类型也是完全不同的
	//for i, v := range b {
	//	fmt.Println(i, v)  // index value
	//}
	//
	//fmt.Println("other pointer")
	//for i := range a {
	//	fmt.Printf("a[%d]: %d\n", i, a[i])
	//}
	//for i, v := range b {
	//	fmt.Printf("b[%d]: %d\n", i, v)
	//}
	//for i := 0; i < len(c); i ++ {
	//	fmt.Printf("c[%d]: %d\n", i, c[i])
	//}

	//var times [5][2][0]int  // 多维数组
	//fmt.Println(times)
	//for range times {
	//	fmt.Println("hello")
	//}

	var s1 = [2]string{"hello", "world"}
	var s2 = [...]string{"你好", "世界"}
	var s2 = [...]string{1: "世界", 0: "你好"}

	// 结构体数组
	var line1 [2]image.Point
	var line2 = [...]image.Point{image.Point{X: 0, Y: 0}, image.Point{X:1, Y:1}}
	var line3 = [...]image.Point{{0, 0}, {1, 1}}

	// 图像解码器数组
	var decoder1 [2]image.Point
	var decoder2 = [...]func(io.Reader) (image.Image, error){
		png.Decode,
		jpeg.Decode,
	}

	// 接口数组
	var unknow1 [1]interface{}
	var unknow2 = [...]interface{}{123, "你好"}

	// 管道数组
	var chanList = [2]chan int{}

	//空数组
	var d [0]int
	var e = [0]int{}
	var f = [...]int{}
}
