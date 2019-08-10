package main

import (
	"fmt"
)

func reverse(str string) string {
	var result string
	strLen := len(str)
	for i := 0; i < strLen; i++ {
		s_u := str[strLen-i-1]
		ss := fmt.Sprintf("%c", s_u)
		fmt.Println(ss, s_u)
		result = result + ss
	}
	return result
}

func reverse1(str string) string {
	tmp := []byte(str)
	length := len(str)
	var result []byte
	for i := 0; i < length; i++ {
		result = append(result, tmp[length-i-1])
	}
	return string(result)
}

func main() {
	var str1 = "hello"
	str2 := "world"

	str3 := str1 + str2
	fmt.Println(str3)

	s3 := fmt.Sprintf("%s %s", str1, str2)
	fmt.Println(s3)

	fmt.Println(s3[:])

	r_str := reverse(s3)
	fmt.Println(r_str)

	r_str = reverse1(s3)
	fmt.Println(r_str)

}
