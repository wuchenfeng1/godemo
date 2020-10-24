package main

import (
	"fmt"
	"time"
)

func main() {
	myChan := make(chan string, 3)
	go count("son", 7, myChan)
	go count("son1", 12, myChan)
	count("main ", 5, myChan)
	for i := 0; i < 3; i++ {
		<-myChan
	}
}

func count(name string, n int, myChan chan<- string) {

	for i := 0; i < n; i++ {
		fmt.Println(name, i)
		time.Sleep(time.Millisecond)
	}
	fmt.Println(name, "  工作完毕")
	myChan <- name + "game over"
}
