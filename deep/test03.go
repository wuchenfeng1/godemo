package main

import "fmt"

//斐波那契数列
func main() {
	for i := 0; i < 100; i++ {
		fmt.Println("", i, "\t", getFbnq(i))
	}

}

func getFbnq(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return getFbnq(n-2) + getFbnq(n-1)

}
