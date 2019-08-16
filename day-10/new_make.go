package main

import "fmt"

/*
new和make都可以用来分配空间，初始化类型，但是他们有一些区别

new
	new(T) 为一个T类型新值分配空间并将此空间初始化为T的零值，返回新值的地址，也就是T类的指针，该指针指向T的x新分配零值

make
	make 只能用于slice, map, channel三种类型，make(T, args)返回的就是初始化后的T类型的值，这个值
	不是T类型的零值，也不是指针*T,是进过初始化后的T的引用
*/

func main () {
	//p1 := new(int)
	//
	//fmt.Printf("p1 ---> %#v \n", p1)
	//fmt.Printf("p1 point to --> %#v \n", *p1)
	//
	//var p2 *int
	//i := 0
	//p2 = &i
	//fmt.Printf("%#v \n", p2)
	//fmt.Printf("%#v \n", *p2)

	// make
	//var s1 []int
	//if s1 == nil {
	//	fmt.Printf("%#v \n", s1)
	//}
	//
	//s2 := make([]int, 1)
	//if s2 == nil {
	//	fmt.Printf("s2 is nil --> %#v \n", s2)
	//
	//} else {
	//	fmt.Printf("s2 --> %#v \n", s2)
	//}

	var m1 map[int]string
	if m1 == nil {
		fmt.Printf("m1 is nil --> %#v \n", m1)
	} else {
		fmt.Printf("m1 not nil --> %#v \n", m1)
	}

	m2 := make(map[int]string)
	if m2 == nil {
		fmt.Printf("m2 is nil --> %#v \n", m2)
	} else {
		fmt.Printf("m2 not nil --> %#v \n", m2)
	}

	var c1 chan string
	if c1 == nil {
		fmt.Printf("c1 is nil --> %#v \n", c1)
	} else {
		fmt.Printf("c1 not nil --> %#v \n", c1)
	}

	var c2 = make(chan string)
	if c2 == nil {
		fmt.Printf("c2 is nil --> %#v \n", c2)
	} else {
		fmt.Printf("c2 not nil --> %#v \n", c2)
	}

	/*
	如果不是特殊声明，go的函数默认都是按照值传参数，即通过传递的参数是是值的副本，在函数内部对值的修改不影响值的本省
	但是make(T, args)返回的值通过函数传递参数之后可以直接修改，即map, slice. channel通过函数传参之后咋函数内部修改将影响函数外部的值

	*/

	s2 := make([]int, 3)
	fmt.Printf("%#v\n", s2)
	modifySlice(s2)
	fmt.Printf("%#v\n", s2)


	type Foo struct {
		name string
		age int
	}

	var foo1 Foo
	fmt.Printf("foo1  %#v\n", foo1)
	foo1.age = 111

	fmt.Println(foo1.age)


	foo3 := &Foo{}
	fmt.Printf("%#v\n", foo3)
	foo3.age = 3
	fmt.Println(foo3.age)

	foo4 := new(Foo)
	fmt.Println(foo4)

	foo4.age = 22
	foo4.name = "aaa"
	fmt.Println(foo4)

	var foo5 *Foo = new(Foo)
	fmt.Printf("%#v\n", foo5)

	foo5.age = 11
	foo5.name = "foo555"
	fmt.Println(foo5)

}

func modifySlice(s []int) {
	s[0]  =1
}


