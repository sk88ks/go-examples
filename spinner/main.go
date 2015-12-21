package main

import (
	"fmt"
	"time"

	"github.com/tj/go-spin"
)

func main() {
	s := spin.New()
	//s.Set(spin.Box7)
	s.Set(spin.Spin9)
	for i := 0; i < 30; i++ {
		fmt.Printf("\r  \033[36mcomputing\033[m %s ", s.Next())
		time.Sleep(100 * time.Millisecond)
	}
}
