package main

import (
	"fmt"

	"github.com/armon/go-radix"
)

func main() {
	// Create a tree
	r := radix.New()

	words := []string{
		"安室奈美恵",
		"LMFAO",
		"This is a test",
		"これはてすとです",
	}

	var idx int
	for i := range words {
		var w string
		for _, s := range words[i] {
			w += string(s)
			fmt.Printf("Insert: %s\n", w)
			r.Insert(w, idx)
			idx++
		}
	}

	m, i, ok := r.Maximum()
	fmt.Println(m, i, ok)
}
