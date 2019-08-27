package main

import (
	"fmt"
	"github.com/pkg/errors"
)

//
//import (
//	"io/ioutil"
//	"github.com/pkg/errors"
//)
//
//func loadConfig()error {
//	_, err:= ioutil.ReadFile("")
//	if err != nil{
//		return errors.Wrap(err, "read failed")
//	}
//}
//
//
//func setup()error{
//	err := loadConfig()
//	if err!=nil{
//		return errors.Wrap(err, "invalid config")
//	}
//}
//
//func foo() (err error) {
//	defer func() {
//		if r:=recover(); r!=nil{
//			switch x := r.(type) {
//			case string:
//				err = errors.New(x)
//			case error:
//				err = x
//			default:
//				err = fmt.Errorf("Unknow panic: %v", r)
//			}
//		}
//	}()
//	panic("TODO")
//}

func main() {
	//foo()
	defer func()(err error) {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = fmt.Errorf("Unknow panic: %v", r)
			}
		}
		fmt.Printf("%v　--> %#v", err, err)
		return err
	}()
	//panic([]chan{})
}
