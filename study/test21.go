package main

import (
	"fmt"
	"time"
)

type jiekou interface {
	new1(num int)
}

type shixian struct {
	port   int
	user   []*user
	mychan chan []*user
}

type user struct {
	name string
	age  int
	addr string
}

func (p *shixian) new1(num int) {
	if num == p.port {
		go demo01(p)
	} else {
		go demo02(p)
	}
	time.Sleep(3 * time.Second)
	fmt.Println("我实现了接口方法")
}

func demo02(p *shixian) {
	x := <-p.mychan

	i := huanxun1(x)
	fmt.Println("demo02", i)
	for _, hehe := range i {
		fmt.Println("哈哈，我向服务器上发送成功了", hehe)
	}

}

func huanxun(users []*user) []*user {
	var bl []*user
	for _, u := range users {
		if u.age == 12 && u.addr == "河南" {
			bl = append(bl, u)
		}
	}
	return bl
}

func huanxun1(users []*user) []*user {
	var bl []*user
	for _, u := range users {
		if u.name == "小米" {
			bl = append(bl, u)
		}
	}
	return bl
}

func demo01(p *shixian) {
	//请求的过滤条件
	i := huanxun(p.user)
	fmt.Println("demo01", i)
	p.mychan <- i
}

func main() {
	var sx = &shixian{}
	users := make(chan []*user, 3)
	sx.user = funcIntDatas()
	sx.mychan = users
	sx.port = 3
	for i := 1; i <= 7; i++ {
		sx.new1(i)
	}

}

func funcIntDatas() []*user {
	var uq []*user
	u := &user{}
	u.name = "小名"
	u.age = 12
	u.addr = "河南"
	uq = append(uq, u)
	u1 := &user{}
	u1.name = "小米"
	u1.age = 12
	u1.addr = "河南"
	uq = append(uq, u1)
	u2 := &user{}
	u2.name = "xiaofeng"
	u2.age = 9
	u2.addr = "河南1"
	u3 := &user{}
	u3.name = "小米"
	u3.age = 12
	u3.addr = "河南"
	uq = append(uq, u3)
	return uq
}
