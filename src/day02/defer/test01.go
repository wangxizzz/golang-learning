package main

import (
	"log"
)

func foo(n int) int {
	log.Println("n1=", n)
	defer func(i int) {
		//n += i
		log.Println("n=", i)
	}(n)
	n += 100
	if n == 200 {
		return n
	}
	log.Println("n2=", n)
	log.Println("函数返回之前执行...")
	return n
}

func main() {
	foo(100)
}
