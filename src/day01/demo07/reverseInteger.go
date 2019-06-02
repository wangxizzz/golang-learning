package main

import "fmt"

func reverse(x int) int {
	result := 0
	for x != 0 {
		b := x % 10
		result = result*10 + b
		x = x / 10
	}
	if result > (1<<31)-1 || result < -1<<31 {
		return 0
	}
	return result
}

func main() {
	fmt.Println(reverse(
		1534236469))
}
