package _continue

import (
	"fmt"
	"testing"
)

/**
测试break跳到标签处。适用于break一次性跳出多层循环
*/
func Test01(t *testing.T) {
	for i := 0; i < 3; i++ {
	next:
		for j := 0; j < 3; j++ {
			fmt.Println(i, j)
			break next
		}
	}
}

/**
测试continue,跳到标签处
*/
func Test02(t *testing.T) {

next:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(i, j)
			continue next
		}
	}

	// 两者的结果不同
	for i := 0; i < 3; i++ {
	next1:
		for j := 0; j < 3; j++ {
			fmt.Println(i, j)
			continue next1
		}
	}
}
