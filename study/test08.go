package main

import "fmt"

//闭包
func main() {
	var zhang, huang int
	task := excuteTask()
	task2 := excuteTask()
	huang = task("黄忠", 12) //将函数的值保存在该闭包中变量中国
	zhang = task2("张飞", 14)
	huang = task("黄忠", 5)
	zhang = task2("张飞", 19)
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
