// Example usage of the Treap data structure.
// This program demonstrates insertion, in-order traversal, search, and memory cleanup
// using the Treap implementation from the treap package.
//
// Author: Braiden Gole
// Created: July 11, 2025
package main

import (
	treap "braidengole/treap/TreapImplementation"
	"fmt"
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
