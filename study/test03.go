package main

import "fmt"

const (
	Monday = iota  // 定义递增表达式
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	fmt.Print(Sunday)
}
