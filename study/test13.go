package main

//对于Go对象

type Car struct {
	color string
	size  string
}

////方式一：使用T{…}方式，结果为值类型
//cc := Car{}
//
////方式二：使用new的方式，结果为指针类型
//c1 := new(Car)
//
////方式三：使用&方式，结果为指针类型
//c2 := &Car{}

//以下为创建并初始化
//c3 := &Car{"红色", "1.2L"}
//c4 := &Car{color: "红色"}
//c5 := Car{color: "红色"}

//构造函数：
//在Go语言中没有构造函数的概念，对象的创建通常交由一个全局的创建函数来完成，以
//NewXXX 来命名，表示“构造函数” ：

func NewCar(color, size string) *Car {
	return &Car{color, size}
}
