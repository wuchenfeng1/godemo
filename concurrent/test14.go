package main

import "sync"

func main() {
	var rw sync.RWMutex
	rw.RLock()   //读锁
	rw.RUnlock() //释放读锁
	rw.Lock()    //写锁
	rw.Unlock()  //释放读锁
}
