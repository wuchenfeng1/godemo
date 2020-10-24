package main

import (
	"fmt"
	"time"
)

func main() {
	go test1()
	//go test2()
	time.Sleep(1 * time.Minute)
	fmt.Println("main")
}

func test1() {
	for i := 0; i < 10000000; i++ {
		fmt.Println("test1")
	}
}
func test2() {
	for i := 0; i < 10; i++ {
		fmt.Println("test2")
	}
}
