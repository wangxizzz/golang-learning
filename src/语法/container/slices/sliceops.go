package main

import "fmt"

func printSlice(s []int) { // 如果加上指针，传地址可以防止每次大量复制数组，消耗性能
	fmt.Printf("%v, len=%d, cap=%d\n",
		s, len(s), cap(s))
}

func modify(s []int) {
	s[0] = 100
}

func sliceOps() {
	fmt.Println("Creating slice")
	var s []int // Zero value for slice is nil。定义一个切片。

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	modify(s1) // 切片改了
	printSlice(s1)

	s2 := make([]int, 16) // 定义定长slice
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]

	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]

	fmt.Println(tail)
	printSlice(s2)
}

func main() {
	sliceOps()
}
