package main

import "fmt"

func main(){
	var a int = 10
	LOOP: for a < 20 {
		if a == 15 {
			a += 1
			goto LOOP
		}
	fmt.Printf("a: %d\n", a)
	a++
	}
	gotoTag()
	
}

func print9x(){
	for m := 1; m < 10; m++ {
		for n := 1; n <= m; n++{
			fmt.Printf("%dx%d=%d", n, m, m * n)
		}
		fmt.Println("")
	}
}

func gotoTag(){
	for m := 1; m < 10; m++ {
		n := 1
		LOOP: if n < m {
			fmt.Println("%dx%d=%d", n, m, m * n)
			n++
			goto LOOP
		} else {
			fmt.Println("")
		}
		n++
	}
}
