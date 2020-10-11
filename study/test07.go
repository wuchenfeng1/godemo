package main

import "fmt"

func main() {
	superStart1("张学友", "刘德华", "黎明", "小明")
}

//不定长的参数
func superStart(args ...string) {
	for i, name := range args {
		fmt.Printf("角标i：%v,name为：%v\n", i, name)

	}

}

//固定参数个数
func superStart1(arg1, arg2, arg3, arg4 string) {
	fmt.Println(arg1, arg2, arg3, arg4)
}
