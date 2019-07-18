// go 切片

package main

import "fmt"

func main() {
	fmt.Println("go　语言切片ss")
	//silce1 := make([]int {1, 2, 3, 4})
	slice2 := []int{1, 2, 3}
	fmt.Println(slice2)

	var numbers = make([]int, 3, 5)
	printSlice(numbers)

	var numbers2 []int

	printSlice(numbers2)
	if (numbers2 == nil) {
		fmt.Printf("切片是空的")
	}

	numbers3 := []int{0, 1, 2, 3, 4, 5, 6}
	printSlice(numbers3)

	/*打印原始切片*/
	fmt.Println("numbers3 == ", numbers3)
	fmt.Println(numbers3[1:4])
	fmt.Println(numbers3[:3])
	fmt.Println(numbers3[4:])
	fmt.Println(numbers3[:])

	numbers = make([]int,0,5)
	printSlice(numbers)

	numbers2 = numbers3[:2]
	printSlice(numbers2)

	numbers4 := numbers3[2:5]
	printSlice(numbers4)

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
