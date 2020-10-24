package main

import (
	"fmt"
	"runtime"
)

func main() {
	var gomaxprocs int
	runtime.NumCPU()                   //查看cpu的核心数
	gomaxprocs = runtime.GOMAXPROCS(2) // 返回先前的cpu的核心数为当前的电脑配置
	fmt.Println(gomaxprocs)
	gomaxprocs = runtime.GOMAXPROCS(1) //2 返回先前的cpu的核心数为2
	fmt.Println(gomaxprocs)
	gomaxprocs = runtime.GOMAXPROCS(3) //1 返回先前的cpu的核心数为1
	fmt.Println(gomaxprocs)

}
