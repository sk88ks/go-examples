package main

import (
	"fmt"

	"github.com/tchap/go-patricia/patricia"
)

func main() {
	printItem := func(prefix patricia.Prefix, item patricia.Item) error {
		fmt.Printf("%q: %v\n", prefix, item)
		return nil
	}

	//// Create a new default trie (using the default parameter values).
	////trie := patricia.NewTrie()

	//// Create a new custom trie.
	trie := patricia.NewTrie(patricia.MaxPrefixPerNode(16), patricia.MaxChildrenPerSparseNode(10))

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
			trie.Insert(patricia.Prefix(w), idx)
			idx++
		}
	}

	//trie.Visit(printItem)

	trie.VisitSubtree(patricia.Prefix("安室"), printItem)

	//// Just check if some things are present in the tree.
	//key := patricia.Prefix("Pepa Novak")
	//fmt.Printf("%q present? %v\n", key, trie.Match(key))
	//// "Pepa Novak" present? true
	//key = patricia.Prefix("Karel")
	//fmt.Printf("Anybody called %q here? %v\n", key, trie.MatchSubtree(key))
	//// Anybody called "Karel" here? true
}
