package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		// handle error
		log.Println(err)
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err1 := bufio.NewReader(conn).ReadString('\n')
	if err1 != nil {
		// handle error
		log.Println(err1)
	}
	log.Println(status)
}
