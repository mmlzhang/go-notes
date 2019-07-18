package main

import "fmt"

func main() {
	var countryCapitalMap map[string]string  // 创建集合
	countryCapitalMap = make(map[string]string)
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Italy１"] = "罗马2"
	countryCapitalMap["Italy１"] = "罗马23"
	countryCapitalMap["Italy２"] = "罗马sdf"

	for countr := range countryCapitalMap {
		fmt.Println(countr, "首都：", countryCapitalMap [countr])
	}

	capital, ok := countryCapitalMap ["Francesss"]
	if (ok) {
		fmt.Println(capital, ok)
	}else {
		fmt.Println(capital, ok, "不存在")
	}
}