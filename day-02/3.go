package main

import "fmt"

func main() {
	var numbers []int
	printSlice(numbers)

	numbers = append(numbers, 0)
	printSlice(numbers)

	numbers = append(numbers, 1)
	printSlice(numbers)

	numbers = append(numbers, 2,3,4)
	printSlice(numbers)

	numbers1 := make([]int, len(numbers), (cap(numbers)) *2)
	copy(numbers1, numbers)
	printSlice(numbers1)

	println("range 范围函数！\n")
	nums := []int{2,3,4}
	sum := 0
	for _, num := range nums {
		sum += num
		println(sum)
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k,v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}


}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
