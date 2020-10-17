package main

import "fmt"

//选择排序 将与第一个数做比较 比较大小，将大的放在左边
func main01() {
	var arr = []int{3, 2, 1, 5, 7, 8, 6, 23, 65, 877, 14}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	fmt.Print(arr)
}

//冒泡排序 原理：比较两个相邻的元素，将值大的元素交换到右边
func main() {
	var arr = []int{3, 2, 1, 5, 7, 8, 6, 23, 65, 877, 14}
	for i := 0; i < len(arr)-1; i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j] > arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	fmt.Print(arr)
}
