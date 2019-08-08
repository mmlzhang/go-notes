package main

import (
	"fmt"
	"unicode/utf8"
)


func str2runes(s []byte) []rune {
	var p []int32
	for len(s) > 0{
		r, size := utf8.DecodeRune(s)
		p = append(p, int32(r))
		s = s[size:]

	}
	return []rune(p)

}

func rune2string(s []int32) string {
	var p []byte
	buf := make([]byte, 3)
	for _, r := range s {
		n := utf8.EncodeRune(buf, r)
		p = append(p, buf[:n]...)

	}
	return string(p)

}

func main() {
	c1 := make(chan [0]int)
	go func() {
		fmt.Println("c1")
		c1 <- [0]int{}
	}()
	<- c1

	c2 := make(chan struct{})
	go func() {
		fmt.Println("c2")
		c2 <- struct{}{}
	}()
	<- c2

	b := [3]int{1, 2, 3}
	fmt.Printf("b: %T\n", b)
	fmt.Printf("b: %#v\n", b)

	fmt.Printf("%#v\n", []byte("Hello, 世界"))
	fmt.Printf("%#v\n", []rune("Hello, 世界"))
	fmt.Printf("%#v\n",  string([]rune{'世', '界'}))
	fmt.Printf("%#v\n",	string([]rune{'世',	'界'}))

	a := "abc"
	b = str2runes(a)
	fmt.Println(b)

}