package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"语法/functional/fib"
)

type intGen func() int

// 表示intGen这个func()函数类型，实现了Reader接口，具体是Read()方法
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	// TODO: incorrect if p is too small!
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	var f intGen = fib.Fibonacci()
	printFileContents(f)
}
