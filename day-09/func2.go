package main

import (
	"net/http/fcgi"
	"image/color"
)

/*
方法一般是面向对象编程的一个特性,在C++语言中方法对饮一个类对象的成员函数,是关联到具体对象上的虚表中的,但是Go语言的方式是
关联到类型的,这样可以在变异阶段完成方法的静态绑定, 一个面向的程序会用方法来表达其属性对应的操作,这样使用这个对象的用户就不需要直接去
操作对象,而是借助方法来做这些事情,面向对象编程进入主流开发领域一般认为是从c++开始的,
*/

type File struct {
	fd int
}

// 打开文件
func OpenFile(name string) (f *File, err error) {
	//
}

// 关闭文件
func CloseFile(f *File) error {
	//
}


//读取数据
func ReadFile(f *File, int64 offset, data []byte) int {
	// ...
}


// 函数和操作对象的类型紧密的绑定在一起
// Go 语言的做法是,讲CloseFile和ReadFile函数的第一个参数移动到函数的开头

// 关闭文件
func (f *File) CloseFile() err {
	//
}

// 读取文件
func (f *File) ReadFile(int64 offset, data []byte) int {
	//
}

// 这样写, CloseFile 和 ReadFile 函数就成了File类型的独有方法了, 他们不在占用包级空间的名字资源
// 同时, File类型已经明确了他们的操作对象,进一步简化

func (f *File) Close() err {
	//
}

func (f *File) Read(int64 offset, data []byte) int {

}

// 还原方法 使用表达式的特还原为普通类型的函数

var CLoseFile = (*File).Close

var ReadFile = (*File).Read

// 文件处理
f, _ := OpenFile("foo.dat")

ReadFile(f, 0, data)
CloseFile(f)


// 简化

// 打开文件
f, _ := OpenFile("foo.dat")

var Close = f.Close
var Read = f.Read

Read(0, data)
Close()


/*
Go 语言不支持传统面向对象中的集成特性,而是有自己特有的组合方式支持了方法的继承,
Go语言中,通过在结构体内置匿名函数的成员来实现继承
*/

import "image/color"

type Point struct{ X, Y float64}

type ColoredPoint struct {
	Point
	Color color.RGBA
}

var cp ColoredPoint

cp.X = 1
fmt.Println(cp.Point.X)

cp.Point.Y = 2
fmt.Println(cp.Y)


