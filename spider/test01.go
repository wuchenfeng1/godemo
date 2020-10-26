package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func main() {
	url := "https://www.kahao.com/index.asp?b_scene_zt=1"
	reponse, err := http.Get(url)
	body := reponse.Body
	defer body.Close()
	handler(err, "解析："+url)
	if reponse.StatusCode == http.StatusOK {
		html, err1 := ioutil.ReadAll(body)
		handler(err1, "解析："+url)
		//fmt.Println(string(all))
		reg := `^1([38][0-9]|14[57]|5[^4])\d{8}$`
		rgx := regexp.MustCompile(reg)
		fmt.Println(rgx.FindAllStringSubmatch(string(html), -1))
	}

}

func handler(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}
}
