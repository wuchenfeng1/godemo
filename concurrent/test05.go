package main

import "fmt"

func main() {
	mychan := make(chan int, 8)
	mychan <- 1000
	fmt.Println(len(mychan)) //1   长度
	fmt.Println(cap(mychan)) //8  容量

}
