package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n",
			id, n)
	}
	// 也可以这样写，但是会打印很多空
	//for {
	//	fmt.Printf("Worker %d received %c\n",
	//		id, <- c)
	//}
	fmt.Println("dsff")

}

func createWorker(id int) chan<- int { // 定义只能往里面发数据的channel
	c := make(chan int)
	// 在goroutin中并发执行work()
	go worker(id, c) //此函数相当于起了一个线程，会立即执行下面一行代码return c.
	// 返回channel
	return c
}

// channel的接收数据只有一个符号 <- ， 只是channel的位置不同。
func chanDemo() {
	var channels [10]chan<- int //定义了10个只能往里面发数据的channel
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}
	// n := <- channels[i] 收数据就会编译报错
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3) // 定义cahnnel，并且设置缓冲区为3
	go worker(0, c)        // 如果只有发数据，没有收数据的channel，就会deadlock
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c) // 表示数据发完了
	time.Sleep(time.Millisecond)
}

func main() {
	//fmt.Println("Channel as first-class citizen")
	//chanDemo()
	//fmt.Println("Buffered channel")
	bufferedChannel()
	//fmt.Println("Channel close and range")
	//channelClose()
}
