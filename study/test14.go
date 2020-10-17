package main

import "fmt"

type Worker interface {
	work(hour int) (product string)
	rest()
}

type Coder struct {
	skill string
}

func (coder *Coder) rest() {
	panic("implement me")
}

func (coder *Coder) work(hour int) (product string) {
	fmt.Printf("程序员工作了%s小时", hour)
	return
}

func main() {
	var worker Worker = &Coder{"122"}
	worker.work(8)

}
