package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	cuurent := make(chan int, 3)
	for i := 0; i < 100; i++ {
		go test10(cuurent, i)
	}

	time.Sleep(1 * time.Minute)

}

func test10(cuurent chan int, i int) {
	cuurent <- 123
	time.Sleep(time.Second)
	fmt.Printf("%d的开方是%.2f \n", i, math.Sqrt(float64(i)))
	<-cuurent
}
