// ===================================================================================
// File:        TreapImplementation.go
// Package:     treap
// Description: This package implements a Treap data structure in Go. A Treap is a
//
//	randomized balanced binary search tree that maintains two invariants:
//	- Binary Search Tree (BST) property based on node keys
//	- Max Heap property based on randomly assigned node priorities
//
//	By combining these properties, the Treap achieves probabilistic
//	balancing and supports efficient insertion, search, and traversal.
//
//	Features implemented in this package:
//	- TreapNode struct with key, priority, left, and right pointers
//	- Randomized priority assignment using a seeded RNG
//	- Recursive insertion with rotations to preserve heap order
//	- Binary search for existing keys
//	- In-order traversal with a callback visitor function
//	- Explicit tree cleanup to release memory (optional in Go)
//
// Author:      Braiden Gole
// Created:     July 17, 2025
//
// Example Usage:
//
//	root := treap.Insert(nil, 10)
//	root = treap.Insert(root, 20)
//	node := treap.Search(root, 20)
//	treap.InOrder(root, func(k, p int) { fmt.Printf("Key: %d, Priority: %d\n", k, p) })
//	treap.Clear(&root)
//
// ===================================================================================
package treapimplementation

import (
	"math/rand"
	"time"
)

// TreapNode represents a node in the Treap.
type TreapNode struct {
	Key      int
	Priority int
	left     *TreapNode
	right    *TreapNode
}

// randomNumberGenerator is used to assign random priorities to nodes.
var randomNumberGenerator *rand.Rand

// init initializes the random number generator with a seed based on current time.
func init() {
	randomNumberGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// rotateLeft performs a left rotation around the given root.
//
//	root            newRoot
//	  \     ==>       /
//	newRoot        root
func rotateLeft(root *TreapNode) *TreapNode {
	var newRoot *TreapNode = root.right

	// Move new root's left subtree to root's right.
	root.right = newRoot.left

	// Place root as left child of new root.
	newRoot.left = root

	return newRoot
}

// rotateRight performs a right rotation around the given root.
//
//	   root            newRoot
//	    /     ==>         \
//	newRoot             root
func rotateRight(root *TreapNode) *TreapNode {
	var newRoot *TreapNode = root.left

	// Move new root's right subtree to root's left.
	root.left = newRoot.right

	// Place root as right child of new root.
	newRoot.right = root

	return newRoot
}

// Insert adds a new key to the Treap while maintaining both BST and heap properties.
// If the key already exists, the Treap remains unchanged.
func Insert(root *TreapNode, key int) *TreapNode {
	if root == nil {
		// Create a new node with a random priority.
		return &TreapNode{
			Key:      key,
			Priority: randomNumberGenerator.Intn(1 << 31),
		}
	}

	if key < root.Key {
		// Recurse into the left subtree.
		root.left = Insert(root.left, key)

		// Heap property violated? Rotate right.
		if root.left != nil && root.left.Priority > root.Priority {
			root = rotateRight(root)
		}
	} else if key > root.Key {
		// Recurse into the right subtree.
		root.right = Insert(root.right, key)

		// Heap property violated? Rotate left.
		if root.right != nil && root.right.Priority > root.Priority {
			root = rotateLeft(root)
		}
	} else {
		// Duplicate key, do nothing...
	}

	return root
}

// Search looks for a key in the Treap and returns the corresponding node.
// Returns nil if the key is not found.
func Search(root *TreapNode, key int) *TreapNode {
	if root == nil || root.Key == key {
		return root
	}

	if key < root.Key {
		return Search(root.left, key)
	}

	return Search(root.right, key)
}

// InOrder performs an in-order traversal of the Treap,
// applying the given visit function to each node's key and priority.
func InOrder(root *TreapNode, visit func(int, int)) {
	if root != nil {
		// Visit the left subtree.
		InOrder(root.left, visit)

		// Visit the current node.
		visit(root.Key, root.Priority)

		// Visit the right subtree.
		InOrder(root.right, visit)
	}
}

// Clears the Treap by recursively setting all node pointers to nil.
// This helps free memory explicitly, although Go's garbage collector handles it.
func Clear(root **TreapNode) {
	if root == nil || *root == nil {
		return
	}

	Clear(&(*root).left)
	Clear(&(*root).right)

	*root = nil
}
