package main

/*
enum C {
	ONE,
	TWO,
};
*/
import "C"
import "fmt"

func main() {
	var c C.enum_C = C.TWO
	fmt.Println(c)
	//fmt.Println(c.ONE)
	fmt.Println(C.ONE)
	fmt.Println(C.TWO)
}