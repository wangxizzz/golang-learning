package main

import (
	_ "goinaction/chapter2/sample/matchers" // 使用空白标识符导入包，避免编译错误。
	"goinaction/chapter2/sample/search"
	"log"
	"os"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	search.Run("president")
}
