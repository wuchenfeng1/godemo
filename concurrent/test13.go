package main

import (
	"fmt"
	"sync"
	"time"
)

//func main() {
//	var wg sync.WaitGroup
//	wg.Add(1) //添加一个协成
//	wg.Done() //注销一个协成
//	wg.Wait() //阻塞等待协成数为零停止等待
//}

func main() {
	var wait sync.WaitGroup
	wait.Add(1)
	go func() {

		fmt.Println("开始1")
		time.Sleep(3 * time.Second)
		fmt.Println("结束1")
		wait.Done()
	}()
	wait.Add(1)
	go func() {
		fmt.Println("开始2")
		timer := time.NewTimer(5 * time.Second)
		<-timer.C
		fmt.Println("结束2")
		wait.Done()
	}()
	wait.Add(1)
	go func() {
		fmt.Println("开始3")
		timer1 := time.NewTicker(1 * time.Second)
		<-timer1.C
		fmt.Println("结束3")
		wait.Done()
	}()
	wait.Wait()
	fmt.Println("over")
}
