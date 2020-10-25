package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var once sync.Once
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			//只执行一次
			once.Do(func() {
				fmt.Println("杀死bi")
			})
			wg.Done()
		}()
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("杀死小米")
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("main over")
}
