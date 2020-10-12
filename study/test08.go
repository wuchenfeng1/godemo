package main

import "fmt"

//闭包
func main() {
	var zhang, huang int
	task := excuteTask01("黄忠")
	task2 := excuteTask01("张飞")
	huang = task(12) //将函数的值保存在该闭包中变量中国
	zhang = task2(14)
	huang = task(5)
	zhang = task2(19)
	fmt.Println(huang)
	fmt.Println(zhang)
}

//闭包函数
func excuteTask() func(name string, count int) (progress int) {
	var progress = 0
	doTask := func(name string, count int) int {
		fmt.Println("%s行军打仗带领s%次", name, count)
		progress += count
		return progress
	}
	return doTask

}

//优化
func excuteTask01(name string) func(count int) (progress int) {
	var progress = 0
	doTask := func(count int) int {
		fmt.Println("%s行军打仗带领s%次", name, count)
		progress += count
		return progress
	}
	return doTask

}
