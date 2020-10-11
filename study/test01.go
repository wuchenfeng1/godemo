package main

import "fmt"

func main() {
	const(
		const01="常量1"
		const02="常量2"
	)//多个常量统一定义

	var(
		var01="变量1"
		var02="变量2"
	)//多个变量量统一定义
	//常量定义
	const pi = 3.14
	//变量定义
	var r = 2.0
	// := 声明+赋值
	are:=pi * r * r
	//打印
	fmt.Println("圆的面积为：",are)
	fmt.Println("多个常量打印：",const01,const02)
	fmt.Println("多个常量打印：",var01,var02)
}
