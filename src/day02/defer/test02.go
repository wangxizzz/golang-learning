package main

import "log"
import "time"
import "sync"

var mu sync.Mutex

// 参考网址：https://www.jianshu.com/p/5b0b36f398a2
func lock() {
	mu.Lock()
	log.Printf("lock")
}

func unlock() {
	mu.Unlock()
	log.Printf("unlock")
}

// defer在函数中间执行
func foo1() int {
	lock()

	func() {
		log.Printf("entry inner")
		defer unlock()
		log.Printf("exit inner")
	}()

	time.Sleep(1 * time.Second)
	log.Printf("return")
	return 0
}

func main() {
	r := foo1()
	log.Println("r=", r)
}
