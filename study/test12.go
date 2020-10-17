package main

import "fmt"

func main() {
	//map的声明   第一种
	var newmap = map[string]int{}
	fmt.Print(newmap)
	//map的赋值
	newmap["233"] = 4
	fmt.Println(newmap)
	var countryCapitalMap = make(map[string]string)
	///* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	bali, ok := countryCapitalMap["France"] // 返回值 和是否存在
	fmt.Println(ok, bali)
	//判断值是都存在，如果
	if newmap, ok := countryCapitalMap["France"]; ok {
		fmt.Println(ok, newmap)
	} else {
		fmt.Println(ok, newmap)
	}
	// 遍历键
	for key := range countryCapitalMap {
		fmt.Println(key)
	}
	// 遍历键，值
	for key, value := range countryCapitalMap {
		fmt.Println(key, value)
	}

}
