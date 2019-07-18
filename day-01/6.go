package main

import "unsafe"

const (
	a = "ab11111111111111c"
	b = len(a)
	c = unsafe.Sizeof(a)

	d = iota + 1
	    o
		f
	i = 1 << iota
	j = 3 << iota
	k
	l
)

func main(){
	println(a, b, c, d, o, f)	
	println(i, j, k, l)
}