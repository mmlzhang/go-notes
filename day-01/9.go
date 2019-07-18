package main

import "fmt"

// 100以内的素数
func main()  {
	var count int
	var flag bool
	count = 1
	for count < 100 {
		count ++
		flag = true
		for tmp := 2; tmp < count; tmp+=1 {
			if count % tmp == 0{
				flag = false
			}
		}
		if flag == true {
			fmt.Println(count, "素数")
		}else{
			continue
		}
	}
}