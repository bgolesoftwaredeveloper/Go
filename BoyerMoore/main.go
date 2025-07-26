// ===================================================================================
// File:        main.go
// Package:     main
// Description: Entry point for demonstrating the Boyer-Moore string search algorithm.
//
//	This file imports the Boyer-Moore implementation and provides an example use case
//	to search for all occurrences of a pattern within a given input text.
//
//	The Boyer-Moore algorithm leverages two preprocessing heuristics:
//	- Bad Character Heuristic: Shifts the pattern based on mismatched characters.
//	- Good Suffix Heuristic: Shifts based on matched suffixes in the pattern.
//
//	Example in this file:
//	- Text:    "XYZXYXZYXYZXYYYXYZXYZZYZX"
//	- Pattern: "XYZ"
//	- Output:  Indices where the pattern is found in the text.
//
// Usage:
//
//	Run this file to see the Boyer-Moore algorithm in action, printing all matched
//	pattern positions based on efficient pattern skipping.
//
// Author:      Braiden Gole
// Created:     July 25, 2025
//
// ===================================================================================
package main

import (
	"fmt"

	boyer_moore "github.com/bgolsoftwaredeveloper/boyer_moore/BoyerMooreImplementation"
)

func main() {
	var text string = "XYZXYXZYXYZXYYYXYZXYZZYZX"
	var pattern string = "XYZ"

	var indices []int = boyer_moore.BoyerMooreSearch(text, pattern)

	fmt.Printf("Pattern '%s' found at indices: %v\n", pattern, indices)
}
