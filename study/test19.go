package main

import (
	"fmt"
	"time"
)

type Man struct {
	Person
}
type Person interface {
	Eat(i int)
}

type Xixi struct {
	name string
	age  int
}

//男人在吃
func (man *Man) Eat(i int) {
	var x = new(Xixi)
	if i == 1 {
		x = hello()
	} else {
		x = haha(x)
	}
	fmt.Println("男人在吃。。。", x)

}

func haha(xixi *Xixi) *Xixi {
	time.Sleep(2 * time.Second)
	xixi.name = "小米"
	fmt.Println("haha。。。", xixi)
	return xixi
}

func hello() (x *Xixi) {
	x = new(Xixi)
	x.age = 12
	time.Sleep(2 * time.Second)
	fmt.Println("hello。。。", x)
	return
}

//人在吃

func main() {
	var p = &Man{}
	for i := 1; i <= 3; i++ {
		p.Eat(i)
	}

}
