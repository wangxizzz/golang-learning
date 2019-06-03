package main

import (
	"fmt"
	"time"
)

func main() {
	//t1 := time.Tick(time.Second)

	d := time.Duration(time.Second * 1)

	t := time.NewTicker(d)
	defer t.Stop()

	go func() {
		for i := 0; i < 3; i++ {
			<-t.C
			fmt.Println("timeout...")
		}
		fmt.Println("bbbbbbbbb")
	}()

	time.Sleep(time.Second * 5)
	fmt.Println("aaaaaaa")
}
