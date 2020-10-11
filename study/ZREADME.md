#够浪~的基本语法练习总结
>常量表达式(代码参考：test01)
 + const pi = 3.14 //常量定义
 + 	const(
        const01="常量1"
        const02="常量2"
    )//多个常量统一定义
 + var(
        var01="变量1"
        var02="变量2"
    )//多个变量量统一定义
 + var r = 2.0 //变量定义
 + are:=pi * r * r // := 声明+赋值
 + fmt.Println("圆的面积为：",are) //打印 与java的打印相同
 + fmt.Printf() //格式化打印 %T 类型 %v 值 %c  字符
>基本数据类型(代码参考：test02.go)
 + *官网的对基本数据类型的解释*
   + 整数的类型官方解释
        ######整数
        + type int8 int8  int8是所有有符号8位整数的集合 Range: -128 through 127.
        + type int16 int16 int16是所有有符号16位整数的集合。范围：-32768到32767。
        + type int int  int是大小至少为32位的有符号整数类型，不是int32的别名
        + type int32 int32 int32是所有有符号32位整数的集合。范围：-2147483648到2147483647。**rune是int32的别名**
        + type int64 int64  int64是所有有符号64位整数的集合。范围：-9223372036854775808到9223372036854775807
        #####
        + type uint8 uint8 uint8是所有无符号8位的集合整数范围：0到255。**type byte = uint8**
        + type uint16 uint16 uint16是所有无符号16位整数的集合。范围：0到65535。
        + type uint uint uint是大小至少为32位的无符号整数类型，不是uint32的别名
        + type uint32 uint32 uint32是所有无符号32位整数的集合。范围：0到4294967295
        + type uint64 uint64 uint64是所有无符号64位整数的集合。范围：0到18446744073709551615
        ######浮点
        + type float32 float32 float32是所有IEEE-754 32位浮点数的集合。
        + type float64 float64 float64 is the set of all IEEE-754 64-bit floating-point numbers.
        ######字符串
        + type string string string是所有8位字节字符串的集合，通常不是这样必须表示UTF-8编码的文本。字符串可以为空，但是 不是零。字符串类型的值是不可变的。
        ##### 布尔
        + type bool bool boolean values, true and false
>基本语法（代码参考：）
+ iota  定义递增表达式 test03.go
+ math  函数  详见test04.go
+ 流程控制  详见test05.go
+ go语言唯一的循环语句 for循环 常见的算法 break、contine、goto、defer  详见test06.go
>函数（代码参考：）
+ 

 




