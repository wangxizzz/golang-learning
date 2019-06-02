package main

import (
	"fmt"
	"testing"
)

func Test01(t *testing.T) {
	fmt.Println("aaa")
}

func Test02(t *testing.T) {
	m := make(map[int]string)
	m[1] = "aa"
	help02(m)
	// 如果这个键存在， map 会返回键所对应值的副本
	fmt.Println(m[1])
}
func help02(m map[int]string) {
	m[1] = "sss"
}
