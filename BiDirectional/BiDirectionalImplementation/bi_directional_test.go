// ===================================================================================
// File:        bi_directional_test.go
// Package:     BiDirectionalImplementation
// Description: This file contains unit tests for the BiDirectionalImplementation
//
//	This package provides a bi-directional tree where each node maintains
//	references to its parent and children.
//
//	The tests in this file verify the correctness and integrity of the
//	core tree operations, including:
//
//	- Relationship maintenance (adding children by value and node reference)
//	- Node search (finding existing and non-existing nodes)
//	- Removing children (valid removals and attempts to remove non-children)
//	- Traversal output (validating PrintUp hierarchical path printing)
//
//	All tests are written using Go’s built-in "testing" package.
//
// Author:      Braiden Gole
// Created:     July 17, 2025
//
// Test Coverage:
//
//	✅ TestAddChildMaintainsRelationship
//	✅ TestAddChildNodeMaintainsRelationship
//	✅ TestFindNodeExists
//	✅ TestFindNodeNotExists
//	✅ TestRemoveChildValid
//	✅ TestRemoveChildInvalid
//	✅ TestPrintUpDisplaysCorrectPath
//
// Usage:
//
//	To run all tests:
//	$ go test -v
//
// ===================================================================================
package BiDirectionalImplementation

import (
	"bytes"
	"os"
	"testing"
)

// ====================
// Relationship Testing
// ====================

// TestAddChildMaintainsRelationship verifies that adding a child by value
// correctly sets the child's parent pointer and updates the parent's children slice.
func TestAddChildMaintainsRelationship(test *testing.T) {
	// Arrange.
	var root *Node = &Node{Value: "Root"}

	// Act.
	var child *Node = root.AddChild("Child")

	// Assert.
	if child.Parent != root {
		test.Error("Expected child parent to be root.")
	}

	if len(root.Children) != 1 || root.Children[0] != child {
		test.Error("Expected root children to contain the new child.")
	}
}

// TestAddChildNodeMaintainsRelationship verifies that adding an existing node as a child
// correctly sets the child's parent pointer and updates the parent's children slice.
func TestAddChildNodeMaintainsRelationship(test *testing.T) {
	// Arrange.
	var root *Node = &Node{Value: "Root"}
	var child *Node = &Node{Value: "Child"}

	// Act.
	root.AddChildNode(child)

	// Assert.
	if child.Parent != root {
		test.Error("Expected child parent to be root.")
	}

	if len(root.Children) != 1 || root.Children[0] != child {
		test.Error("Expected root children to include attached child.")
	}
}

// ============
// Find Testing
// ============

// TestFindNodeExists verifies that searching for an existing node by value
// returns the correct node pointer.
func TestFindNodeExists(test *testing.T) {
	// Arrange.
	var root *Node = &Node{Value: "Root"}
	var child *Node = root.AddChild("Child")
	var grandchild *Node = child.AddChild("Grandchild")

	// Act.
	var found *Node = root.Find("Grandchild")

	// Assert.
	if found != grandchild {
		test.Errorf("Expected to find grandchild node got, %v.", found)
	}
}

// TestFindNodeNotExists verifies that searching for a non-existent node value
// returns nil.
func TestFindNodeNotExists(test *testing.T) {
	// Arrange.
	var root *Node = &Node{Value: "Root"}

	root.AddChild("Child")

	// Act.
	var found *Node = root.Find("Missing")

	// Assert.
	if found != nil {
		test.Errorf("Expected nil when searching for non-existent node, got %v.", found)
	}
}

// ==============
// Remove Testing
// ==============

// TestRemoveChildValid verifies that removing a valid child node updates
// the parent's children slice and sets the child's parent pointer to nil.
func TestRemoveChildValid(test *testing.T) {
	// Arrange.
	var root *Node = &Node{Value: "Root"}
	var child *Node = root.AddChild("Child")

	// Act.
	var removed bool = root.RemoveChild(child)

	// Assert.
	if !removed {
		test.Error("Expected child to be removed.")
	}

	if child.Parent != nil {
		test.Error("Expected child's parent to be nil after removal.")
	}

	if len(root.Children) != 0 {
		test.Error("Expected root to have 0 children after removal.")
	}
}

// TestRemoveChildInvalid verifies that attempting to remove a node that
// is not a child returns false and does not alter the tree.
func TestRemoveChildInvalid(test *testing.T) {
	// Arrange.
	var root *Node = &Node{Value: "Root"}
	var stranger *Node = &Node{Value: "Stranger"}

	// Act.
	var removed bool = root.RemoveChild(stranger)

	// Assert.
	if removed {
		test.Error("Expected false when removing node not in children.")
	}
}

// =================
// Traversal Testing
// =================

// captureOutput is a helper function that captures and returns
// standard output produced by the provided function.
func captureOutput(function func()) string {
	readPipe, writePipe, _ := os.Pipe()

	// Save original stdout.
	var original *os.File = os.Stdout

	os.Stdout = writePipe

	// Run the function that writes to stdout.
	function()

	// Restore stdout and close writer.
	_ = writePipe.Close()

	os.Stdout = original

	// Read from the read pipe.
	var buffer bytes.Buffer

	// Read from the read end of the pipe into the buffer.
	_, _ = buffer.ReadFrom(readPipe)

	// Return the collected output as a string.
	return buffer.String()
}

func TestPrintUpDisplaysCorrectPath(test *testing.T) {
	// Arrange.
	var root *Node = &Node{Value: "Root"}

	var child *Node = root.AddChild("Child")
	var grandchild *Node = child.AddChild("Grandchild")

	var expected string = "Grandchild <- Child <- Root\n"

	// Act.
	var output string = captureOutput(func() {
		grandchild.PrintUp()
	})

	// Assert.
	if output != expected {
		test.Errorf("Expected PrintUp output %q, got %q.", expected, output)
	}
}
