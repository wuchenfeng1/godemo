package main

import (
	"fmt"
	"time"
)

func main() {
	//定时器 一次性的
	timer := time.NewTimer(2 * time.Second)
	fmt.Println(time.Now())
	//停止防止计时器启动。
	//如果调用停止计时器，则返回true；如果计时器已经停止，则返回false
	//过期或被停止
	//timer.Stop()//停止
	timer.Reset(time.Second) //重置
	x := <-timer.C
	fmt.Println(x)
	//秒 周期性
	ticker := time.NewTicker(2 * time.Second)

	for {
		fmt.Println(<-ticker.C)
	}

}
