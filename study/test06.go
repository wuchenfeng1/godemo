package main

import "fmt"

//死循环
func main04() {
	for true {
		fmt.Println("来了")
	}
}

//for循环
func main05() {
	sum := 0
	for i := 1; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}

//for循环  类似while
func main06() {
	sum := 0
	for sum <= 10 {
		sum += sum
		sum++
		fmt.Print(sum)
	}
	fmt.Println(sum)
}

//foreach 与python for in 写法相似
func main07() {
	str := []string{"google", "runoob"}
	for i, s := range str {
		fmt.Println(i, s)
	}
}

//foreach 与python for in 写法相似
func main08() {
	str := []string{"google", "runoob"}
	for i, s := range str {
		fmt.Println(i, s)
	}
}

//以下实例使用循环嵌套来输出 2 到 100 间的素数：
func main09() {
	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break // 如果发现因子，则不是素数
			}
		}
		if j > (i / j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
}

//九九乘法表：
func main10() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%d ", i, j, j*i)
		}
		fmt.Println("")
	}
}

/**
goto 就是一种标记 告诉你跳过哪一行直接过去直接哪一行代码
*/
func main12() {
	goto START
	fmt.Println("你猜猜")
START:
	fmt.Println("你在哪")
	return
END:
	fmt.Println("我在这个")
	goto END
}

/**
defer 延迟加载 函数
*/
func main() {
	defer main12()
	fmt.Println("我先执行")
	//匿名函数的编写
	defer func() {
		fmt.Println("我是匿名函数")
	}()
}
