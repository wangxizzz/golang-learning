package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	fmt.Println(time.Now())
	select {
	case e1 := <-ch1:
		//如果ch1通道成功读取数据，则执行该case处理语句
		fmt.Printf("1th case is selected. e1=%v", e1)
	case e2 := <-ch2:
		//如果ch2通道成功读取数据，则执行该case处理语句
		fmt.Printf("2th case is selected. e2=%v", e2)
	case a := <-time.After(2 * time.Second):
		// 超时后，会返回当前时间
		fmt.Println("Timed out", a)
	}
}
