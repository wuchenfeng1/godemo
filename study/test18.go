package main

import (
	"log"
	"strconv"
	"time"
)

type Produce struct {
	name string
}

func main() {
	//商品管道
	chanShop := make(chan Produce, 5)
	////技术管道
	//chanCount := make(chan int, 5)
	chuShop := make(chan Produce, 3)
	//生产者源源不断的向[商品管道]写入茶品
	go Producer(chanShop)
	//消费者源源不断的向[商品管道]读入茶品
	go Consumer(chanShop, chuShop)
	x := <-chuShop

	log.Println("give over", x)

}

func Consumer(chanShop <-chan Produce, chuShop chan Produce) {
	//for  {
	l := <-chanShop
	l.name = "呵呵"
	log.Println("消费者吃掉了：", l.name)
	time.Sleep(3 * time.Second)
	chuShop <- l
	//}
}

func Producer(chanShop chan<- Produce) {
	//for  {
	produce := Produce{name: "产品：" + strconv.Itoa(time.Now().Second())}
	log.Println("生产者：", produce)
	time.Sleep(5 * time.Second)
	chanShop <- produce

	//}

}
