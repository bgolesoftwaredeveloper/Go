// ===================================================================================
// File:        treap_test.go
// Package:     treap
// Description: This file contains unit tests for the Treap implementation. A Treap
//
//	is a randomized balanced binary search tree that maintains both BST
//	ordering and heap-based priority balancing.
//
//	The tests in this file verify the correctness and integrity of the
//	Treap's core operations, including:
//
//	- Left and right rotations (rotation correctness and property retention)
//	- Insertions (duplicate handling, heap ordering, and in-order key ordering)
//	- Searches (positive, negative, root, and empty-treap cases)
//	- Memory cleanup via explicit clearing of the treap
//
//	All tests are written using Go’s built-in "testing" package.
//
// Author:      Braiden Gole
// Created:     July 11, 2025
//
// Test Coverage:
//
//	✅ TestRotateLeftPreservesProperties
//	✅ TestRotateRightPreservesProperties
//	✅ TestInsertKeysPresent
//	✅ TestInsertMaintainsBSTProperty
//	✅ TestInsertMaintainsHeapProperty
//	✅ TestInsertIgnoresDuplicates
//	✅ TestInsertIntoEmpty
//	✅ TestInsertWithRotationsMaintainsProperties
//	✅ TestSearchFound
//	✅ TestSearchNotFound
//	✅ TestSearchEmptyTreap
//	✅ TestSearchRootKey
//	✅ TestClearEmptiesTreap
//
// Usage:
//
//	To run all tests:
//	$ go test
//
// ===================================================================================
package TreapImplementation

import "testing"

// ================
// Rotation Testing
// ================

// TestRotateLeftPreservesProperties ensures that a left rotation
// properly updates the treap structure while keeping BST and heap properties.
func TestRotateLeftPreservesProperties(test *testing.T) {
	// Arrange.
	var root *TreapNode = &TreapNode{Key: 10, Priority: 10}

	var rightChild *TreapNode = &TreapNode{Key: 20, Priority: 20}

	root.right = rightChild

	// Act.
	var newRoot *TreapNode = rotateLeft(root)

	// Assert.
	if newRoot != rightChild {
		test.Error("Expected new root to be right child.")
	}

	if newRoot.left != root {
		test.Error("Expected original root to be left child of new root.")
	}

	// Additional BST property check.
	if newRoot.left.Key >= newRoot.Key {
		test.Errorf("BST property violated: left child key %d >= new root key %d.", newRoot.left.Key, newRoot.Key)
	}

	// Additional heap property check.
	if newRoot.Priority < newRoot.left.Priority {
		test.Errorf("Heap property violated: root priority %d < left child priority %d.", newRoot.Priority,
			newRoot.left.Priority)
	}
}

// TestRotateRightPreservesProperties ensures that a right rotation
// properly updates the treap structure while keeping BST and heap properties.
func TestRotateRightPreservesProperties(test *testing.T) {
	// Arrange.
	var root *TreapNode = &TreapNode{Key: 20, Priority: 10}

	var leftChild *TreapNode = &TreapNode{Key: 10, Priority: 20}

	root.left = leftChild

	// Act.
	var newRoot *TreapNode = rotateRight(root)

	// Assert.
	if newRoot != leftChild {
		test.Error("Expected new root to be left child.")
	}

	if newRoot.right != root {
		test.Error("Expeced original root to be right child of new root.")
	}

	// Additional BST property check.
	if newRoot.right.Key <= newRoot.Key {
		test.Errorf("BST property violated: right child key %d <= new root key %d.", newRoot.right.Key, newRoot.Key)
	}

	// Additional heap property check.
	if newRoot.Priority < newRoot.right.Priority {
		test.Errorf("Heap property violated: root priority %d < right child priority %d.", newRoot.Priority, newRoot.right.Priority)
	}
}

// ==============
// Insert Testing
// ==============

// TestInsertKeysPresent verifies that all inserted keys can be found using Search.
func TestInsertKeysPresent(test *testing.T) {
	// Arrange.
	var root *TreapNode

	// Insert some keys.
	var keys []int = []int{50, 30, 70, 20, 40, 60, 80}

	// Act.
	for _, key := range keys {
		root = Insert(root, key)
	}

	// Assert.
	for _, key := range keys {
		if Search(root, key) == nil {
			test.Errorf("Key %d not found after insertion.", key)
		}
	}
}

// TestInsertMaintainsBSTProperty ensures that in-order traversal
// of the treap produces a sorted sequence of keys.
func TestInsertMaintainsBSTProperty(test *testing.T) {
	// Arrange.
	var root *TreapNode

	var keys []int = []int{50, 30, 79, 20, 40, 60, 80}

	// Act.
	for _, key := range keys {
		root = Insert(root, key)
	}

	var inOrderKeys []int

	InOrder(root, func(key int, priority int) {
		inOrderKeys = append(inOrderKeys, key)
	})

	// Assert.
	for index := 1; index < len(inOrderKeys); index++ {
		if inOrderKeys[index-1] >= inOrderKeys[index] {
			test.Errorf("In-order traversal not sorted: %v", inOrderKeys)
		}
	}
}

// TestInsertMaintainsHeapProperty checks that every node's priority
// is greater than or equal to its children's priorities, maintaining heap order.
func TestInsertMaintainsHeapProperty(test *testing.T) {
	// Arrange.
	var root *TreapNode

	var keys []int = []int{50, 30, 70, 20, 40, 60, 80}

	// Act.
	for _, key := range keys {
		root = Insert(root, key)
	}

	// Assert.
	var checkHeapProperty func(node *TreapNode) bool

	checkHeapProperty = func(node *TreapNode) bool {
		if node == nil {
			return true
		}

		if node.left != nil && node.left.Priority > node.Priority {
			return false
		}

		if node.right != nil && node.right.Priority > node.Priority {
			return false
		}

		return checkHeapProperty(node.left) && checkHeapProperty(node.right)
	}

	if !checkHeapProperty(root) {
		test.Error("Heap property violated.")
	}
}

// TestInsertIgnoresDuplicates verifies that inserting a duplicate key
// does not modify the treap.
func TestInsertIgnoresDuplicates(test *testing.T) {
	// Arrange.
	var root *TreapNode

	var keys []int = []int{50, 30, 70}

	for _, key := range keys {
		root = Insert(root, key)
	}

	// Act.
	var rootBefore *TreapNode = root

	// Duplicate key.
	root = Insert(root, 30)

	// Assert.
	if root != rootBefore {
		test.Error("Treap change after inserting duplicate key.")
	}
}

// TestInsertIntoEmpty tests that inserting into an empty treap
// creates a root node with the expected key.
func TestInsertIntoEmpty(test *testing.T) {
	// Arrange.
	var root *TreapNode

	// Act.
	root = Insert(root, 100)

	// Assert.
	if root == nil {
		test.Fatal("Insert failed: root is nil.")
	}

	if root.Key != 100 {
		test.Errorf("Expected root key 100, got %d.", root.Key)
	}
}

// TestInsertWithRotationsMaintainsProperties verifies that after inserting
// multiple keys (which may cause rotations), the treap maintains both
// the binary search tree (BST) ordering and the heap property with respect
// to priorities.
func TestInsertWithRotationsMaintainsProperties(test *testing.T) {
	// Arrange.
	var root *TreapNode

	var keys []int = []int{50, 30, 70, 20, 40, 60, 80}

	for _, key := range keys {
		root = Insert(root, key)
	}

	// Act.
	var inOrderKeys []int

	InOrder(root, func(key int, priority int) {
		inOrderKeys = append(inOrderKeys, key)
	})

	for index := 1; index < len(inOrderKeys); index++ {
		if inOrderKeys[index-1] >= inOrderKeys[index] {
			test.Errorf("BST property violated after insert with rotations: %v.", inOrderKeys)
		}
	}

	// Assert.
	var checkHeap func(node *TreapNode) bool

	checkHeap = func(node *TreapNode) bool {
		if node == nil {
			return true
		}

		if node.left != nil && node.left.Priority > node.Priority {
			return false
		}

		if node.right != nil && node.right.Priority > node.Priority {
			return false
		}

		return checkHeap(node.left) && checkHeap(node.right)
	}

	if !checkHeap(root) {
		test.Error("Heap property violated after insert with rotations.")
	}
}

// ==============
// Search Testing
// ==============

// TestSearchFound ensures that existing keys are found by Search.
func TestSearchFound(test *testing.T) {
	// Arrange.
	var root *TreapNode

	var keys []int = []int{50, 30, 70, 20, 40, 60, 80}

	for _, key := range keys {
		root = Insert(root, key)
	}

	// Act & Assert.
	for _, key := range keys {
		var node *TreapNode = Search(root, key)

		if node == nil {
			test.Errorf("Expected to find key %d, but got nil", key)
		} else if node.Key != key {
			test.Errorf("Found node key %d, expected %d", node.Key, key)
		}
	}
}

// TestSearchNotFound ensures Search returns nil for keys not in the treap.
func TestSearchNotFound(test *testing.T) {
	// Arrange.
	var root *TreapNode

	var keys []int = []int{50, 30, 70}

	for _, key := range keys {
		root = Insert(root, key)
	}

	// Act.

	// Key not present.
	var node *TreapNode = Search(root, 999)

	// Assert
	if node != nil {
		test.Errorf("Expected nil for key 999, but got node with key %d", node.Key)
	}
}

// TestSearchEmptyTreap ensures Search returns nil if treap is empty.
func TestSearchEmptyTreap(test *testing.T) {
	// Arrange.
	var root *TreapNode

	// Act.
	var node *TreapNode = Search(root, 10)

	// Assert.
	if node != nil {
		test.Errorf("Expected nil when searching empty treap, got %v", node)
	}
}

// TestSearchRootKey ensures Search finds the root key correctly.
func TestSearchRootKey(test *testing.T) {
	// Arrange.
	var root *TreapNode

	root = Insert(root, 42)

	// Act.
	var node *TreapNode = Search(root, 42)

	// Assert.
	if node == nil || node.Key != 42 {
		test.Errorf("Expected to find root key 42, got %v", node)
	}
}

// =============
// Clear Testing
// =============

func TestClearEmptiesTreap(test *testing.T) {
	// Arrange.
	var root *TreapNode

	var keys []int = []int{50, 30, 70}

	for _, key := range keys {
		root = Insert(root, key)
	}

	// Act.
	Clear(&root)

	// Assert.
	if root != nil {
		test.Error("Expected root to be nil after clear.")
	}
}
