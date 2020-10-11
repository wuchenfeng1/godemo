package main

import (
	"fmt"
	"math"
)

func main() {
    //四舍五入
	round := math.Round(3.14)
	fmt.Println(round)
	round1 := math.Round(-3.14)
	fmt.Println(round1)
	//绝对值
	fmt.Println(math.Abs(-1231))
	//平方
	fmt.Println(math.Pow(-2,3))
	//开方
	fmt.Println(math.Sqrt(8))
	//开立方
	fmt.Println(math.Round(math.Pow(8,1.0/3)))
	// 正弦
	fmt.Println(math.Sin(30.0/180)*math.Pi)
	//余弦
	fmt.Println(math.Cos(30.0/180)*math.Pi)
	//正切
	fmt.Println(math.Tan(30.0/180)*math.Pi)


}
