package main

import (
	"fmt"
	"sync"
	"time"
)

//数据库 一读多多写
//创建40个协成  30个读取数据库   10个写  并发数控制在5
func main() {
	var wait sync.WaitGroup
	var rws sync.RWMutex
	ints := make(chan int, 5)
	for i := 0; i < 40; i++ {
		wait.Add(1)
		go readDabase(&rws, &wait, ints)
	}
	wait.Add(1)
	for i := 0; i < 10; i++ {
		go writeDabase(&rws, &wait, ints)
	}

	wait.Wait()
	fmt.Println("main 结束2")
}
func readDabase(rws *sync.RWMutex, wait *sync.WaitGroup, ints chan int) {
	ints <- 123
	rws.RLock()
	fmt.Println("readDabase开始")
	<-time.After(time.Second)
	fmt.Println("readDabase结束2")
	rws.RUnlock()
	wait.Done()
	<-ints
}
func writeDabase(rws *sync.RWMutex, wait *sync.WaitGroup, ints chan int) {
	ints <- 123
	rws.Lock()
	fmt.Println("writeDabase开始")
	<-time.After(time.Second)
	fmt.Println("writeDabase结束2")
	rws.Unlock()
	wait.Done()
	<-ints
}
