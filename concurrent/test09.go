package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
生产者每秒生成一件商品，并通知物流公司取货
物流公司将商品运输到超市
消费者阻塞等待从商品消费
消费10轮就主协成结束
*/
func main() {
	logisticsChan := make(chan *produce, 1)
	chaoshiChan := make(chan *produce, 1)
	//整点的时候结束任务
	zhengdianChan := make(chan int, 1)
	go producer(logisticsChan, zhengdianChan)
	go logisticsCompany(logisticsChan, chaoshiChan)
	go consumer(chaoshiChan)

	<-zhengdianChan
	//time.Sleep(2*time.Minute)

}

//消费者阻塞等待从商品消费
func consumer(chaoshiChan <-chan *produce) {
	for {
		p := <-chaoshiChan
		fmt.Println("消费者了一个产品：", p)

	}

}

//物流公司将商品运输到超市
func logisticsCompany(logisticsChan <-chan *produce, chaoshiChan chan<- *produce) {
	for {
		x := <-logisticsChan
		chaoshiChan <- x
	}
}

//生产者每秒生成一件商品，并通知物流公司取货
func producer(logisticsChan chan<- *produce, zhengdianChan chan int) {
	for {
		p := &produce{name: "产品" + strconv.Itoa(time.Now().Second())}
		fmt.Println("生产者生产了一个产品：", p)
		time.Sleep(time.Second)
		logisticsChan <- p
		if time.Now().Second() == 0 {
			zhengdianChan <- 123
		}

	}
}

type produce struct {
	name string
}
