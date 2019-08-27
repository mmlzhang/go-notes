package main

import (
	"log"
	"syscall"
	"fmt"
	"os"
	"io"
	"html"
)

/*
错误处理和异常

排除异常的情况，如果程序运行失败仅被认为是几个预期的结果之一，对于那些讲运行失败看做是预期结果的函数，
他们会返回一个额外的值，通常最后一个来传递错误信息，如果导致失败的原因只有一个，额外的返回值可以是一个布尔值，
通常被命名为ok,
 比如从一个map查询一个结果时，可以通过额外的布尔值判断是否成功
　if v, ok := m["key"]; ok {
      return v
 }

在Go语言中，错误被认为是一种可以预期的结果；而异常则是一种非预期的结果，发生异常可能表示程序中存在BUG或发生了其他
不可控制的问题，Go语言推荐使用recover函数讲内部异常转为错误处理，这使得用户可以真正关心业务相关的　错误处理

如果某个借口简单地讲所有普通的错误当做异常抛出，将会使错误信息杂乱且没有价值，就像在main函数中直接捕获一样，是没有意义的

*/

func main() {
	err := syscall.Chmod(":invalid path", 0666)
	//if err != nil {
	//	fmt.Println(err)
	//	log.Fatal(err.(syscall.Errno))
	//}

	//在main函数中直接捕获异常
	defer func() {
		if r:=recover();r !=nil{
			log.Fatal(r)
		}
	}()

	//if _, err == html.Parse(resp.Body); err != nil {
	//	return nil,  fmt.Errorf("")
	//}
}

// 文件复制，错误处理
func CopyFile(dstName, srcName string)(written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}

	dst, err := os.Create(dstName)
	if err != nil{
		return
	}
	written, err = io.Copy(dst, src)
	defer dst.Close()
	defer src.Close()
	return
}

// recover使用，讲异常转化为普通错误
func ParseJSON(input string) (s *Syntax, err error) {
	defer func() {
		if p:=recover();p!=nil {
			err = fmt.Errorf("JSON: interal error: %#v", p)
		}
	}()
	return
}

