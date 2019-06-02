package main

import (
	"fmt"
	"time"
)

func main() {
	// 多个goroutin对应1个线程
	for i := 0; i < 1000; i++ {
		// 开启1000个线程去执行下面的函数。fun()是一个匿名函数，可以把fun(j)函数体提到外面定义
		go func(j int) {
			for {
				fmt.Printf("Hello from "+
					"goroutine %d\n", j)
			}
		}(i)
	}
	time.Sleep(time.Minute)
}
