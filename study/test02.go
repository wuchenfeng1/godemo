package main

import "fmt"

/**
   go基本数据类型有四种整数 浮点  布尔 字符串
 */
func main() {
	//官网的对基本数据类型的解释
	//整形
	//type int8 int8  int8是所有有符号8位整数的集合 Range: -128 through 127.
	//type int16 int16 int16是所有有符号16位整数的集合。范围：-32768到32767。
	//type int int // int是大小至少为32位的有符号整数类型，不是int32的别名
	//type int32 int32 int32是所有有符号32位整数的集合。//范围：-2147483648到2147483647。rune是int32的别名
	//type int64 int64  int64是所有有符号64位整数的集合。范围：-9223372036854775808到9223372036854775807

	//type uint8 uint8 uint8是所有无符号8位的集合整数范围：0到255。 type byte = uint8
	//type uint16 uint16 uint16是所有无符号16位整数的集合。范围：0到65535。
	//type uint uint uint是大小至少为32位的无符号整数类型，不是uint32的别名
	//type uint32 uint32 uint32是所有无符号32位整数的集合。范围：0到4294967295
	//type uint64 uint64 uint64是所有无符号64位整数的集合。范围：0到18446744073709551615


	//浮点类型

	//type float32 float32 float32是所有IEEE-754 32位浮点数的集合。
	//
	//type float64 float64 float64 is the set of all IEEE-754 64-bit floating-point numbers.

	//字符串

	//type string string string是所有8位字节字符串的集合，通常不是这样必须表示UTF-8编码的文本。字符串可以为空，但是 不是零。字符串类型的值是不可变的。
	//布尔
	//type bool bool 值为true或者false
	var v1 int=100
	var v2 float32=32
	var v3 string="乔"
	var v4 bool=true
    // 格式化打印 %T 类型 %v 值 %c  字符
	fmt.Printf("v1 %T:",v1)
	fmt.Printf("v2 %v:",v2)
	fmt.Printf("v3 %c:",v3)
	fmt.Printf("v4 %c:",v4)

}
