package main

import "fmt"

//函数的运用
func main() {
	//superStart1("张学友", "刘德华", "黎明", "小明")
	i := sum(4, 7)
	he := sum2(3, 7)
	h, cha, ji, shang := jishu(8, 2)
	fmt.Println(i)
	fmt.Println(he)
	fmt.Println(h)
	fmt.Println(cha)
	fmt.Println(ji)
	fmt.Println(shang)
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

//返回值
func sum(a, b int) int {
	return a + b
}

//自定义 返回名称返回
func sum2(a, b int) (he int) {
	he = a + b
	return
}

// 多返回值
func jishu(a, b int) (he, cha, ji, shang int) {
	he = a + b
	cha = a - b
	ji = a * b
	shang = a / b
	return
}
