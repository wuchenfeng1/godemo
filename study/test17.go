package main

import (
	"encoding/json"
	"fmt"
)

type Stu struct {
	Name string
	Age  int
}

func main() {
	stu := Stu{Name: "12", Age: 12}
	//jsonStu, err:= json.Marshal(p)
	jsonStu, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}
	fmt.Println(string(jsonStu))
}
