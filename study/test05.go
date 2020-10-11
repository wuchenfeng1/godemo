package main

import (
	fmt "fmt"
	"time"
)

/**
if else if  else
*/
func main01() {
	fmt.Println("请输入你喜欢的明星：")
	var name string
	//接受用户输入
	fmt.Scan(&name)
	if name == "张靓颖" {
		fmt.Println("我也喜欢很有才华")
	} else if name == "杨幂" {
		fmt.Println("这个女的长的还行")
	} else {
		fmt.Println("我也喜欢明星")
	}
}

/**
switch case  default
*/
func main02() {
	fmt.Println("请输入你喜欢的明星：")
	var name string
	//接受用户输入
	fmt.Scan(&name)
	switch name {
	case "张靓颖":
		fmt.Println("我也喜欢张靓颖")
	case "杨幂":
		fmt.Println("我也喜欢杨幂")
	case "张艺兴":
		fmt.Println("我也喜欢张艺兴")
	default:
		fmt.Println("我也喜欢明星")
	}
}

/**
案例1 如果有一个或多个IO操作可以完成，
则Go运行时系统会随机的选择一个执行，
否则的话，如果有default分支，
则执行default分支语句，如果连default都没有，
则select语句会一直阻塞，直到至少有一个IO操作可以进行
*/
func main() {
	start := time.Now()
	c := make(chan interface{})
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(4 * time.Second)
		close(c)
	}()
	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 3
	}()
	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- 5
	}()
	fmt.Println("Blocking on read...")
	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	case <-ch1:
		fmt.Printf("ch1 case...")
	case <-ch2:
		fmt.Printf("ch1 case...")
	default:
		fmt.Printf("default go...")
	}
}
