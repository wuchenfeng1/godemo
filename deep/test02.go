package main

import "fmt"

func main() {
	fmt.Print(digui(10))
	fmt.Print(sumNum(10))
}

func sumNum(n int) (sums int) {
	for i := 1; i <= n; i++ {
		sums += i
	}
	return

}

//递归加减法
func digui(i int) (sum int) {
	if i == 1 {
		return 1
	}
	return i + digui(i-1)
}
