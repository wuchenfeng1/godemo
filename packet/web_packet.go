package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	//body := bytes.NewReader(data)
	body1 := strings.NewReader("groupName=alerting,alertName,请求超时信息告警")
	response, _ := http.Post("http://127.0.0.1:8090/api/test1", "application/json", body1)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}
