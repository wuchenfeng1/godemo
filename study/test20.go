package main

import "fmt"

type Humaner interface {
	SayHi()
}

type Student struct {
	name string
	id   int
}

func (tmp *Student) SayHi() {
	fmt.Printf("Student%s %d sayhi\n", tmp.name, tmp.id)
}

type Teacher struct {
	group string
	addr  string
}

func (tmp *Teacher) SayHi() {
	fmt.Printf("Teacher%s %s sayhi\n", tmp.group, tmp.addr)
}

type MyStr string

func (tmp *MyStr) SayHi() {
	fmt.Printf("MyStr %s sayhi", *tmp)
}

func main() {

	//只是实现了此接口方法的类型，那么这个类型的变量（接收者类型）就可以给i赋值
	var humaner Humaner = &Student{"mike", 1}
	humaner.SayHi()
}
