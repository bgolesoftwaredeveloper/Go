// ===================================================================================
// File:        aho_corasick.go
// Package:     aho
// Description: This package implements the Aho-Corasick string matching algorithm
//
//	The Aho-Corasick algorithm is a multi-pattern string matching algorithm
//	that builds a finite state machine to search for multiple patterns
//	simultaneously in linear time.
//
//	Features implemented in this package:
//	- Trie-based pattern insertion
//	- Failure link construction (like KMP fallback logic)
//	- Efficient multi-pattern search with overlap support
//
//	The algorithm is useful in applications such as virus scanning,
//	natural language processing, lexical analysis, and intrusion detection.
//
// Author:      Braiden Gole
// Created:     July 19, 2025
//
// ===================================================================================
package ahocorasickimplementation

import "container/list"

// Node represents a single state in the Aho-Corasick trie.
// Each node maintains links to child nodes, a failure link, and a list of patterns matched at that node.
type Node struct {
	children map[rune]*Node
	fail     *Node
	output   []string
}

// AhoCorasick represents the main automaton structure containing the root node.
type AhoCorasick struct {
	root *Node
}

// NewAhoCorasick initializes and returns a new instance of the Aho-Corasick automaton.
func NewAhoCorasick() *AhoCorasick {
	return &AhoCorasick{
		root: &Node{
			children: make(map[rune]*Node),
			output:   []string{},
		},
	}
}

// AddPattern inserts a pattern into the trie, character by character.
// Each new character creates a new node if it doesn't already exist.
func (aho *AhoCorasick) AddPattern(pattern string) {
	var node *Node = aho.root

	// Traverse (or build) down the trie based on each rune in the pattern.
	for _, character := range pattern {
		if _, exists := node.children[character]; !exists {
			// Create a new node if the path doesn't exist.
			node.children[character] = &Node{
				children: make(map[rune]*Node),
				output:   []string{},
			}
		}

		node = node.children[character]
	}

	// Register the complete pattern at the terminal node.
	node.output = append(node.output, pattern)
}

// BuildFailureLinks constructs the failure links (fallbacks) used during pattern search.
// This step is essential to enable fast traversal when mismatches occur.
func (aho *AhoCorasick) BuildFailureLinks() {
	var queue *list.List = list.New()

	// Set fail links of depth-1 children to root and enqueue them for BFS.
	for _, child := range aho.root.children {
		child.fail = aho.root
		queue.PushBack(child)
	}

	var current *Node = nil
	var fail *Node = nil

	// BFS traversal to build failure links for all nodes.
	for queue.Len() > 0 {
		// Dequeue the current node.
		current = queue.Remove(queue.Front()).(*Node)

		// Process all children of the current node.
		for character, child := range current.children {
			fail = current.fail

			// Start following failure links from the parent’s fail node.
			for fail != nil && fail.children[character] == nil {
				fail = fail.fail
			}

			// Set the fail link for the child node.
			if fail == nil {
				child.fail = aho.root
			} else {
				// Inherit the matched child.
				child.fail = fail.children[character]

				// Merge output patterns from the fail node.
				child.output = append(child.output, child.fail.output...)
			}

			// Enqueue the child for further BFS traversal.
			queue.PushBack(child)
		}
	}
}

// Search scans the given text for all patterns previously added to the trie.
// Returns a map from matched pattern to list of starting indices in the text.
func (aho *AhoCorasick) Search(text string) map[string][]int {
	var result map[string][]int = make(map[string][]int)

	var node *Node = aho.root

	// Iterate through each rune in the input text.
	for index, character := range text {
		// Follow failure links if no match.
		for node != aho.root && node.children[character] == nil {
			node = node.fail
		}

		// Transition to next state if possible.
		if next, exists := node.children[character]; exists {
			node = next
		}

		// Record all matched patterns at this node.
		for _, pattern := range node.output {
			result[pattern] = append(result[pattern], index-len(pattern)+1)
		}
	}

	return result
}
