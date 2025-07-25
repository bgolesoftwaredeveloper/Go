// ===================================================================================
// File:        boyer_moore_test.go
// Package:     boyermooreimplementation
// Description: This file contains unit tests for the Boyer-Moore string search
//
//	algorithm implementation.
//
// The tests cover multiple scenarios to verify the correctness of the
// BoyerMooreSearch function, including:
//   - Exact matches at various positions
//   - Multiple occurrences of patterns
//   - Cases with no matches
//   - Edge cases like empty patterns or patterns longer than text
//   - Support for Unicode characters and overlapping patterns
//
// Author:      Braiden Gole
// Created:     July 25, 2025
//
// ===================================================================================
package boyermooreimplementation

import (
	"testing"
)

// equalIntSlices compares two integer slices for equality.
// Returns true if both slices are the same length and contain
// the same elements in order; otherwise returns false.
func equalIntSlices(compare []int, against []int) bool {
	if len(compare) != len(against) {
		return false
	}

	for index := range compare {
		if compare[index] != against[index] {
			return false
		}
	}

	return true
}

// TestBoyerMoorSearch runs a set of table-driven tests for BoyerMooreSearch.
// It verifies the function correctly finds all occurrences of a pattern
// within a given text string. Test cases include edge conditions and
// typical usage scenarios, checking returned indices against expected results.
func TestBoyerMoorSearch(test *testing.T) {
	var tests = []struct {
		name     string
		text     string
		pattern  string
		expected []int
	}{
		{
			name:     "Example test",
			text:     "hello",
			pattern:  "ll",
			expected: []int{2},
		},
		{
			name:     "Exact match in middle",
			text:     "say hello to the world",
			pattern:  "hello",
			expected: []int{4},
		},
		{
			name:     "Multiple occurrences",
			text:     "abracadabra",
			pattern:  "abra",
			expected: []int{0, 7},
		},
		{
			name:     "No match",
			text:     "abcdefg",
			pattern:  "xyz",
			expected: []int{},
		},
		{
			name:     "Pattern equals text",
			text:     "pattern",
			pattern:  "pattern",
			expected: []int{0},
		},
		{
			name:     "Empty pattern",
			text:     "nonempty",
			pattern:  "",
			expected: []int{},
		},
		{
			name:     "Pattern longer than text",
			text:     "short",
			pattern:  "longpattern",
			expected: []int{},
		},
		{
			name:     "Unicode characters",
			text:     "日本語のテキストとパターン",
			pattern:  "テキスト",
			expected: []int{4},
		},
		{
			name:     "Overlapping patterns",
			text:     "aaaaa",
			pattern:  "aaa",
			expected: []int{0, 1, 2},
		},
	}

	for _, specificTest := range tests {
		test.Run(specificTest.name, func(individualTest *testing.T) {
			var result []int = BoyerMooreSearch(specificTest.text, specificTest.pattern)

			if !equalIntSlices(result, specificTest.expected) {
				individualTest.Errorf("BoyerMooreSearch(%q, %q) = %v; want %v", specificTest.text, specificTest.pattern,
					result, specificTest.expected)
			}
		})
	}
}
