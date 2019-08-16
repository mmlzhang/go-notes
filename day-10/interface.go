package main

import "fmt"

type Itface interface {
	Get() int
	Set(int)
}

type S struct {
	Age int
}


func (s S) Get() int {
	return s.Age
}


func (s *S) Set(age int) {
	s.Age = age
}

func f(i Itface) {
	i.Set(10)
	fmt.Println(i.Get())
}


func main(){

	/*
	interface 的变量中存储的是实现了 interface 的类型的对象值，
	这种能力是 duck typing。在使用 interface 时不需要显式在 struct 上声明要实现
	哪个 interface ，只需要实现对应 interface 中的方法即可，go 会自动进行 interface 的检查，
	并在运行时执行从其他类型到 interface 的自动转换，即使实现了多个 interface，go
	也会在使用对应 interface 时实现自动转换，这就是 interface 的魔力所在。
	*/
	s := S{}
	//f(&s)

	var i Itface
	i = &s
	s.Set(100)
	fmt.Println(i.Get())

	if t, ok := i.(*S); ok {
		fmt.Println(t, ok)
	}

	switch t := i.(type) {
	case *S:
		fmt.Println("store *S", t)
	case *R:
		fmt.Println("store *R", t)
	}

	/*
	空的interface  interface{}
	空的interface表示没有方法，所以可以认为所有的类型都实现了该interface{}
	如果定义一个函数参数是interface{} 类型，这个函数可以接受任何类型的参数
	*/

}
