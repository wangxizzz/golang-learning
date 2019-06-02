// Sample program to show how a bytes.Buffer can also be used
// with the io.Copy function.
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// main is the entry point for the application.
func main() {
	var b bytes.Buffer

	// Write a string to the buffer.
	b.Write([]byte("Hello"))

	// 使用 Fprintf 将字符串拼接到 Buffer。b的指针类型实现了Writer接口，所以实例化Writer接口，需要传递引用。
	// Use Fprintf to concatenate a string to the Buffer.
	fmt.Fprintf(&b, "World!")

	// Write the content of the Buffer to stdout.
	io.Copy(os.Stdout, &b)
}
