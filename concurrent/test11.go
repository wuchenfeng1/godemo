package main

import (
	"fmt"
	"time"
)

func main() {
	chaA := make(chan int, 5)
	chaB := make(chan int, 4)
	chaB <- 123
	chaB <- 123
	chaB <- 123
	chaB <- 123
	chaC := make(chan int, 2)
	chaC <- 123
	chaC <- 123

flag: //跳出标记位置
	for {
		select {
		case chaA <- 12:
			fmt.Println("随机执行任务,选择了我A")
			time.Sleep(time.Second)
		case <-chaB:
			fmt.Println("随机执行任务,选择了我B")
			time.Sleep(time.Second)
		case <-chaC:
			fmt.Println("随机执行任务,选择了我C")
			time.Sleep(time.Second)
		default:
			fmt.Println("没有任务,情况下我来执行    *^-^*   ")
			time.Sleep(time.Second)
			break flag //跳出标记位置
		}
	}
	fmt.Println("任务结束 over")
}
