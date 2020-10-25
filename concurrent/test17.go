package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	randchan := make(chan int, 1)
	wg.Add(1)
	go a(randchan, &wg)
	wg.Add(1)
	go b(randchan, &wg)
	go c()
	wg.Wait()
	fmt.Println("over......")
}

func c() {
	fmt.Println("我啥也没干")
}

func getRandom(start int, end int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return start + r.Intn(end-start+1)
}

func b(random chan int, s *sync.WaitGroup) {

	//for  {
	rand := <-random
	if getJO(rand) {
		fmt.Println("收到的随机数为：", rand, "为偶数")
	} else {
		fmt.Println("收到的随机数为：", rand, "为奇数")
	}

	//}
	s.Done()
}

func getJO(r int) bool {
	return r%2 == 0
}

func a(rand chan int, s *sync.WaitGroup) {
	ticker := time.NewTicker(time.Second)
	//for  {
	random := getRandom(100, 999)
	fmt.Println("生成随机数", random)
	<-ticker.C
	rand <- random

	//}
	s.Done()
	fmt.Println(time.Now())

}
