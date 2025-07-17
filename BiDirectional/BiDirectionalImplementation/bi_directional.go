// Package bi_directional provides a basic implementation of a bi-directional tree,
// where each node maintains references to its parent and children.
// This example demonstrates adding, finding, removing nodes, and printing the tree.
//
// Author: Braiden Gole
// Created: July 17, 2025
package BiDirectionalImplementation

import "fmt"

// Node represents a node in a bi-directional tree.
// Each node has a value, a pointer to its parent, and a slice of children.
type Node struct {
	Value    string
	Parent   *Node
	Children []*Node
}

// AddChild creates a new child node with the given value, attaches it to the current node,
// and returns a pointer to the newly created child node.
func (node *Node) AddChild(value string) *Node {
	child := &Node{
		Value:  value,
		Parent: node,
	}

	node.Children = append(node.Children, child)

	return child
}

// AddChildNode attaches an existing node as a child of the current node.
// It sets the child's Parent pointer to the current node and appends the child to the Children slice.
func (node *Node) AddChildNode(child *Node) {
	child.Parent = node
	node.Children = append(node.Children, child)
}

// Find searches the tree recursively starting from the current node
// for a node containing the specified value. Returns a pointer to the found node or nil if not found.
func (node *Node) Find(value string) *Node {
	if node.Value == value {
		return node
	}

	for _, child := range node.Children {
		if result := child.Find(value); result != nil {
			return result
		}
	}

	return nil
}

// RemoveChild removes the specified child node from the current node's children slice.
// It also sets the removed child’s parent pointer to nil. Returns true if the child was found and removed.
func (node *Node) RemoveChild(child *Node) bool {
	for index, descendant := range node.Children {
		if descendant == child {
			// Remove from slice.
			node.Children = append(node.Children[:index], node.Children[index+1])
			child.Parent = nil

			return true
		}
	}

	return false
}

// PrintDown prints the tree structure starting from the current node down to all descendants.
// The level argument is used to control indentation for hierarchical display.
func (node *Node) PrintDown(level int) {
	prefix := ""

	for index := 0; index < level; index++ {
		prefix += "    "
	}

	fmt.Println(prefix + node.Value)

	for _, child := range node.Children {
		child.PrintDown(level + 1)
	}
}

// PrintUp prints the path from the current node up to the root of the tree.
// Each node value is printed in order from leaf to root.
func (node *Node) PrintUp() {
	current := node

	for current != nil {
		fmt.Print(current.Value)

		if current.Parent != nil {
			fmt.Print(" <- ")
		}

		current = current.Parent
	}

	fmt.Println()
}
