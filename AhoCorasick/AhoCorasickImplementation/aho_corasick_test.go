// ===================================================================================
// File:        aho_corasick_test.go
// Package:     AhoCorasickImplementation
// Description: This file contains unit tests for the AhoCorasickImplementation
//
//	This package implements the Aho-Corasick string matching algorithm,
//	efficiently searching multiple patterns in a single pass over the text.
//
//	The tests in this file verify the correctness and robustness of the
//	algorithm through a range of scenarios, including:
//
//	- Adding single and multiple patterns
//	- Building failure links for fallback transitions
//	- Performing pattern matching over varied input strings
//	- Handling empty inputs and overlapping matches
//
//	All tests use Go’s standard "testing" package.
//
// Author:      Braiden Gole
// Created:     July 19, 2025
//
// Test Coverage:
//
//	✅ TestAddPatternAndSearchSingle         — Single pattern matching in text
//	✅ TestMultiplePatterns                  — Matching multiple patterns with overlaps
//	✅ TestNoMatches                         — Ensures no false positives in unmatched text
//	✅ TestOverlappingPatterns               — Tests behavior with nested and overlapping patterns
//	✅ TestEmptyPatternsAndText              — Verifies handling of empty pattern and text edge cases
//
// Usage:
//
//	To run all tests:
//	$ go test -v
//
// ===================================================================================
package ahocorasickimplementation

import (
	"reflect"
	"testing"
)

// TestAddPatternAndSearchSingle tests matching a single pattern ("he") in a basic text input.
func TestAddPatternAndSearchSingle(test *testing.T) {
	// Arrange.
	var ahoCorasick *AhoCorasick = NewAhoCorasick()

	var pattern string = "he"

	ahoCorasick.AddPattern(pattern)

	ahoCorasick.BuildFailureLinks()

	var text string = "hello there"

	var expected map[string][]int = map[string][]int{
		"he": {0, 7},
	}

	// Act.
	var result map[string][]int = ahoCorasick.Search(text)

	// Assert.
	if !reflect.DeepEqual(result, expected) {
		test.Errorf("Search() = %v; want %v.", result, expected)
	}
}

// TestMultiplePatterns tests multiple pattern detection within a single string.
func TestMultiplePatterns(test *testing.T) {
	// Arrange.
	var ahoCorasick *AhoCorasick = NewAhoCorasick()

	var patterns []string = []string{"he", "she", "his", "hers"}

	for _, pattern := range patterns {
		ahoCorasick.AddPattern(pattern)
	}

	ahoCorasick.BuildFailureLinks()

	var text string = "ushers"

	var expected map[string][]int = map[string][]int{
		"she":  {1},
		"he":   {2},
		"hers": {2},
	}

	// Act.
	var result map[string][]int = ahoCorasick.Search(text)

	// Assert.
	if !reflect.DeepEqual(result, expected) {
		test.Errorf("Search() = %v; want %v.", result, expected)
	}
}

// TestNoMatches verifies that no patterns match when the text does not contain them.
func TestNoMatches(test *testing.T) {
	// Arrange.
	var ahoCorasick *AhoCorasick = NewAhoCorasick()

	ahoCorasick.AddPattern("abc")

	ahoCorasick.BuildFailureLinks()

	var text string = "defghijkl"

	var expected map[string][]int = map[string][]int{}

	// Act.
	var result map[string][]int = ahoCorasick.Search(text)

	// Assert.
	if !reflect.DeepEqual(result, expected) {
		test.Errorf("Search() = %v; want %v.", result, expected)
	}
}

// TestOverlappingPatterns tests behavior when patterns overlap and are nested within each other.
func TestOverlappingPatterns(test *testing.T) {
	// Arrange.
	var ahoCorasick *AhoCorasick = NewAhoCorasick()

	var patterns []string = []string{"a", "ab", "bab", "bc", "bca", "c", "caa"}

	for _, pattern := range patterns {
		ahoCorasick.AddPattern(pattern)
	}

	ahoCorasick.BuildFailureLinks()

	var text string = "abccab"

	var expected map[string][]int = map[string][]int{
		"a":  {0, 4},
		"ab": {0, 4},
		"bc": {1},
		"c":  {2, 3},
	}

	// Act.
	var result map[string][]int = ahoCorasick.Search(text)

	// Assert.
	if !reflect.DeepEqual(result, expected) {
		test.Errorf("Search() = %v; want %v.", result, expected)
	}
}

// TestEmptyPatternsAndText checks handling of empty input text and patterns.
func TestEmptyPatternsAndText(test *testing.T) {
	// Arrange.
	var ahoCorasick *AhoCorasick = NewAhoCorasick()

	// No patterns added, search empty text.
	ahoCorasick.BuildFailureLinks()

	// Act.
	var result map[string][]int = ahoCorasick.Search("")

	// Assert.
	if len(result) != 0 {
		test.Errorf("Search() on empty input = %v; want empty map.", result)
	}

	// Arrange.

	// Add empty pattern and search empty text.
	ahoCorasick.AddPattern("")

	ahoCorasick.BuildFailureLinks()

	// Act.
	result = ahoCorasick.Search("")

	// Assert.
	if len(result) != 0 {
		test.Errorf("Search() with empty pattern = %v; want empty map.", result)
	}
}
