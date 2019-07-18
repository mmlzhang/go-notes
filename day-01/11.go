package main

func main(){
	a := max(11, 22)
	println(a)
}

func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	}else{
		result = num2
	}
	return result
}