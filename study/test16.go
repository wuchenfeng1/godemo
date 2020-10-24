package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	file, err := ioutil.ReadFile("D:\\111\\1.txt")
	if err != nil {
		fmt.Println("文件出现异常", err.Error())
	}
	fmt.Println(string(file))

}
func main1() {
	file, err := os.Open("D:\\111\\1.txt")
	if err != nil {
		fmt.Println("文件出现异常", err.Error())
	}
	defer func() {
		file.Close()
		fmt.Println("文件关闭....")
	}()
	reader := bufio.NewReader(file)
	for true {
		readString, readErr := reader.ReadString('\n')
		if readErr != nil || readErr == io.EOF {
			fmt.Println("文件读取完毕！！！！")
			break
		} else {
			fmt.Println("文件出现异常", err.Error())
			return
		}
		fmt.Println(readString)
	}

}
