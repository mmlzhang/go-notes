package main

import "fmt"


/*
空的interface{}表示可以接受任何类型的值
*/
func testempty(i interface{}) {
	fmt.Println(i)
}

/*
空的interface slice   []interface{}  不能接受所有类型的slice，不能自动转换
由于　interface{} 会占用两个字节的存储空间，一个是自身的methods数据，一个是指向其存储值的指针，也就是interface的变量
因此　[]interface{} 的古典长度是 N*2  但是　[]T　的固定长度是　N*sizeof(T) 两种slice的实际存储值的大小不同
我们可以自己自行转换

var dataSlice []int = foo()
var interfaceSlice []interface{} = make([]interface{}, len(dataSlice))
for i, d := range dataSlice {
    interfaceSlice[i] = d
}
*/


// receiver 是 value，函数用 pointer 的形式调用

type I interface {
	Get() int
	Set(int)
}


type SS struct {
	Age int
}

func (s SS) Get() int {
	return s.Age
}

func (s SS) Set(age int) {
	s.Age = age
}


func ff (i I) {
	i.Set(10)
	fmt.Println(i.Get())
}


func main (){
	ss := SS{}

	ff(&ss)
	ff(ss)
}

//func main() {
//	var i string = "aaaa"
//	testempty(i)
//}
