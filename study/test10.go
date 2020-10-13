package main

import "fmt"

func main() {
	////声明长度为10的一个数组
	//var balance [10] int
	////声明长度为10的一个数组
	//var balance1 [...] int
	//定义一个长度为6的数组
	var arr1 = [6]int{12, 123, 12, 123, 123, 123}
	//定义一个长度未知的数组
	var arr = [...]int{12, 123, 12, 123, 123, 123, 232, 23, 23, 23, 23, 23, 23}
	//给数组赋值
	//balance[0]=1223
	//fmt.Print("数组arr的长度：",len(balance))
	//fmt.Print("数组arr的长度：",len(balance1))
	fmt.Print("数组arr的长度：", len(arr1))
	fmt.Print("数组arr的长度：", len(arr))
	fmt.Printf("数组的类型%T,%v", arr, arr)
}
