// ===================================================================================
// File:        main.go
// Description: Example usage of the Treap data structure implemented in Go. A Treap
//
//	combines the properties of a binary search tree (BST) and a heap,
//	maintaining both the BST ordering by keys and the heap ordering by
//	randomly assigned priorities.
//
//	This program demonstrates the following Treap operations:
//	- Insertion of integer keys (with random priorities)
//	- In-order traversal to display nodes in sorted order
//	- Search for specific keys
//	- Memory cleanup via explicit tree clearing (for completeness)
//
//	The Treap implementation is imported from the "treap" package and
//	used here to showcase practical applications of probabilistic trees
//	in maintaining balanced binary search structures.
//
// Author:      Braiden Gole
// Created:     July 17, 2025
//
// Usage:
//
//	To run the program:
//	$ go run main.go
//
// ===================================================================================
package main

import (
	"fmt"

	treap "github.com/bgolesoftwaredeveloper/treap/TreapImplementation"
)

func main() {
	var root *treap.TreapNode

	// Insert keys.
	root = treap.Insert(root, 50)
	root = treap.Insert(root, 30)
	root = treap.Insert(root, 70)
	root = treap.Insert(root, 20)
	root = treap.Insert(root, 40)
	root = treap.Insert(root, 60)
	root = treap.Insert(root, 80)

	// Perform in-order traversal (sorted by key).
	fmt.Println("In-order traversal:")

	treap.InOrder(root, func(key int, priority int) {
		fmt.Printf("Key %d, Priority: %d\n", key, priority)
	})

	fmt.Println()

	// Search for a key.
	var key int = 40

	var node *treap.TreapNode = treap.Search(root, key)

	if node != nil {
		fmt.Printf("Found key %d with priority %d\n", node.Key, node.Priority)
	} else {
		fmt.Printf("Key %d not found in treap.\n", key)
	}

	// Free the treap (not strictly necessary in Go, but included for completeness).
	treap.Clear(&root)
}
