package main

import "fmt"

func main(){
	var a bool = false
	var b bool = true
	if ( a && b ) {
		fmt.Printf("第一行　－　条件为　true\n")
	}
	if ( a || b ) {
		fmt.Printf("条件为false")
	}
	if ( !a ) {
		fmt.Printf("!否定")
	}
}
