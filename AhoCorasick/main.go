// ===================================================================================
// File:        main.go
// Description: Example usage of the Aho-Corasick multi-pattern string matching algorithm
//
//	The Aho-Corasick algorithm efficiently searches multiple string patterns
//	simultaneously within a given text by constructing a trie with failure links,
//	enabling linear-time scanning.
//
//	This program demonstrates key operations of the Aho-Corasick implementation:
//	- Adding multiple string patterns to the trie
//	- Building failure links to enable fallback during mismatches
//	- Searching a sample input text for all occurrences of the patterns
//	- Printing all matched patterns with their corresponding start indices
//
//	This example showcases practical applications of the Aho-Corasick algorithm
//	in areas such as text processing, intrusion detection, and lexical analysis.
//
// Author:      Braiden Gole
// Created:     July 19, 2025
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

	aho_corasick "github.com/bgolesoftwaredeveloper/aho_corasick/AhoCorasickImplementation"
)

func main() {

	var ahoCorasick *aho_corasick.AhoCorasick = aho_corasick.NewAhoCorasick()

	// Define the set of patterns to search for.
	var patterns []string = []string{"he", "she", "his", "hers"}

	// Add each pattern to the automaton's trie structure.
	for _, pattern := range patterns {
		ahoCorasick.AddPattern(pattern)
	}

	// Construct failure links for efficient fallback during search.
	ahoCorasick.BuildFailureLinks()

	// Define the input text to scan.
	var text string = "ushers"

	// Perform the search across the input text.
	var result map[string][]int = ahoCorasick.Search(text)

	fmt.Println("Matches found:")

	// Iterate through matched patterns and their positions.
	for pattern, positions := range result {
		for _, index := range positions {
			fmt.Printf("\tPattern '%s' found at index %d\n", pattern, index)
		}
	}
}
