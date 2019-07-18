package main

import "fmt"

func main() {

	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	fmt.Println("原始地图")

	for country := range countryCapitalMap {
		fmt.Println(country, "首都：－", countryCapitalMap [country])
	}

	// 删除
	delete(countryCapitalMap, "France")
	fmt.Println("删除发过后的地图")

	for country := range countryCapitalMap {
		fmt.Println(country, "首都", countryCapitalMap [country])
	}
}
