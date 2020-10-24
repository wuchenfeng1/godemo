package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//runtime.Gosched()
	for i := 1; i <= 3; i++ {
		go func(num int) {
			for j := 1; j <= 3; j++ {
				if num == 3 {
					runtime.Gosched()
				}
				fmt.Println("协成:", num, "的值为:", j)
			}
		}(i)
	}
	fmt.Println("122")
	fmt.Println("aaa")
	time.Sleep(3 * time.Second)
	fmt.Println("ccccc")
}
