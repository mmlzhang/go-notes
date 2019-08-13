package main

import (
	"testing"
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

//
//var (
//	a io.ReadCloser = (*os.File)(f) //隐式转换, *os.File满足 io.ReadCloser接口
//	b io.Reader = a // 隐式转换, io.ReadClose满足 io.Reader 接口
//	c io.Closer = a  // 隐式转换  io.ReadCloser 满足 io.Closer接口
//	d io.Reader = c.(io.Reader)  // 显示转换, io.CLoser 不满足 io.Reader 接口
//)

type TB struct {
	testing.TB
}

func (p *TB) Fatal(args ...interface{}) {
	fmt.Println("TB.Fatal disabled!")
}

func main() {
	var tb testing.TB = new(TB)
	tb.Fatal("Hello, playground!")
}


// 模拟实现一个grpc插件

type grpcPlugin struct {
	*generator.Generator
}

func (p *grpcPlugin) Name() string {return "grpc"}

func (p *grpcPlugin) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *grpcPlugin) GenerateImports(file *generator.FileDescriptor)  {
	if len(file.Service) == 0 {
		reutrn
	}
	p.P(`import "google.golang.org/grpc"`)
}

