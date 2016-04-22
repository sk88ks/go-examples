package main

import (
	"fmt"

	"github.com/willf/bloom"
)

func main() {
	var n uint
	n = 1000
	filter := bloom.New(20*n, 5)
	filter.Add([]byte("Test"))

	if filter.Test([]byte("Test")) {
		fmt.Println("Hit")
	}
}
