package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var chengguan = false
	//变量协成
	cond := sync.NewCond(&sync.Mutex{})
	dageMsg := make(chan string, 1)

	for i := 0; i < 3; i++ {
		go func(num int) {
			for {
				cond.L.Lock()
				if chengguan {
					select {
					case dageMsg <- "城管来了。。。":
						fmt.Println(num, "城管来了。。。")
						cond.Wait() //等待
					default:

					}
				}
				cond.L.Unlock()
				fmt.Println(num, "提供露腿活动。。。")
				time.Sleep(3 * time.Second)

			}
		}(i)
	}

	go func() {
		for {
			select {
			case x := <-dageMsg:
				cond.L.Lock()
				time.Sleep(3 * time.Second)
				chengguan = false
				cond.Signal() //唤醒所有
				fmt.Println(x, "摆平。。。")
				cond.L.Unlock()
			default:

				fmt.Println("公关专员的幸福生活。。。")
				time.Sleep(3 * time.Second)
				chengguan = true
			}

		}

	}()

	time.Sleep(365 * time.Second)
	fmt.Println("main over")
}
