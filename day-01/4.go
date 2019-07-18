package main

var x, y int

var (  // 这种因式分解关键字写法一般用于申明全局变量
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

func numbers()(int, int, string){
	a, b, c := 1, 2, "str"
	return a, b, c
}
func main(){
	g, h := 111, "world"
	println(x, y, a, b, c, d ,e, f, g, h)

	_, numb, strs := numbers()
	println(numb, strs)
}