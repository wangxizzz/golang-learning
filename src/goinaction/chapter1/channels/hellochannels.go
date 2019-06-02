package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printer(ch chan int) {
	// for循环会一直阻塞，直到有数据来了。在调用了close(),for循环就会退出，否则就会陷入deadlock
	for i := range ch {
		fmt.Printf("Received %d ", i)
	}
	fmt.Println("接受完毕....")
	wg.Done()
}

// main is the entry point for the program.
func main() {
	c := make(chan int)
	go printer(c)
	wg.Add(1)

	// Send 10 integers on the channel.
	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c) // close(c)需要在wait()之前调用，因为需要先退出for循环，才能执行wg.Done().
	wg.Wait()

}
