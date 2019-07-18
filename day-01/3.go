package main

import "fmt"

func main()  {
	/*var intVar int
	intVar := 1*/
	indVar := 1
	fmt.Println(indVar)

	var v1, v2, v3 int
	v1, v2, v3 = 1, 2, 3
	fmt.Println(v1, v2, v3)

	var vv1, vv2, vv3 = 1, 2, 3
	fmt.Println(vv1, vv2, vv3)

	a1, a2, a3 := 1, 2, 3
	fmt.Println(a1, a2, a3)
}

var (
	v1 int
	v2 int
)