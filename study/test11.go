package main

import "fmt"

func main() {
	//切片的类型是 []int 数组的类型是[10]int
	//切片与python的数组角标的思想一直
	var arr = [...]int{12, 123, 12, 123, 123, 123, 232, 23, 23, 23, 23, 23, 23, 23223233}
	// 切片
	ints := arr[0:9]
	//切片
	ints1 := arr[:9]
	//数组
	ints2 := arr[2:]
	//数组转切片
	ints3 := arr[:]

	//切片 可以在切片
	i := ints3[:]
	fmt.Println(ints)
	fmt.Println(ints1)
	fmt.Println(ints2)
	fmt.Println(ints3)
	fmt.Println(i)
}
