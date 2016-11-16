package main

import (
	"fmt"
	"unicode/utf8"

	"github.com/armon/go-radix"
	"github.com/k0kubun/pp"
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
			i, ok := r.Insert(w, idx)
			fmt.Println(i, ok)
			idx++
		}
	}

	for i := range words {
		var w string
		for _, s := range words[i] {
			w += string(s)
			fmt.Printf("Insert: %s\n", w)
			i, ok := r.Insert(w, nil)
			fmt.Println(i, ok)
			idx++
		}
	}

	var currentV string
	var currentL int
	var res []string
	r.Walk(func(k string, v interface{}) bool {
		l := utf8.RuneCountInString(string(k))
		if l < currentL {
			res = append(res, currentV)
		}
		fmt.Println(k)
		currentL = l
		currentV = k
		//fmt.Println(currentV)

		return false
	})
	res = append(res, currentV)

	pp.Println(res, len(res))

	m, i, ok := r.Maximum()
	fmt.Println(m, i, ok)
}
