// ===================================================================================
// File:        boyer_moore.go
// Package:     boyermooreimplementation
// Description: This package implements the Boyer-Moore string search algorithm in Go.
//
//	Boyer-Moore is an efficient pattern matching algorithm that performs
//	substring search in O(n) time on average by preprocessing the pattern
//	using two heuristics:
//
//	- Bad Character Heuristic: Skips alignments based on mismatched characters.
//	- Good Suffix Heuristic: Shifts the pattern using known suffix matches.
//
//	Features implemented in this package:
//	- Full preprocessing of bad character and good suffix tables
//	- Maximum shift selection per iteration for optimal skipping
//	- Returns all starting indices of pattern occurrences in the input text
//
// Author:      Braiden Gole
// Created:     July 25, 2025
//
// ===================================================================================
package boyermooreimplementation

// maximum returns the greater of two integer values.
// Used to determine the optimal shift between the bad character and good suffix heuristics.
func maximum(compare int, against int) int {
	if compare > against {
		return compare
	}

	return against
}

// preprocessBadCharacterTable creates the bad character table for the given pattern.
// The table maps each rune in the pattern to its last occurrence index.
// This table helps determine how far to shift the pattern when a mismatch occurs.
func preprocessBadCharacterTable(patternRunes []rune) map[rune]int {
	var table map[rune]int = make(map[rune]int)

	for index, character := range patternRunes {
		table[character] = index
	}

	return table
}

// preprocessSuffixes computes the suffix lengths array for the pattern.
// For each position, it calculates the length of the longest suffix
// of the substring pattern[0:index] that matches a suffix of the entire pattern.
// This array is used in constructing the good suffix shift table.
func preprocessSuffixes(pattern []rune) []int {
	var patternLength = len(pattern)
	var suffixLengths []int = make([]int, patternLength)

	// The entire pattern is always suffix of itself.
	suffixLengths[patternLength-1] = patternLength

	var searchRightBoundary int = patternLength - 1
	var searchLeftBoundary int = 0

	// Process suffix lengths from right to left (excluding last character).
	for index := patternLength - 2; index >= 0; index-- {
		// Case 1: Reuse previously computed suffix length if within the right boundary.
		if index > searchRightBoundary && suffixLengths[index+patternLength-1-searchLeftBoundary] < index-searchRightBoundary {
			suffixLengths[index] = suffixLengths[index+patternLength-1-searchLeftBoundary]
		} else {
			// Case 2: Calculate suffix length by matching characters from the right.
			if index < searchRightBoundary {
				searchRightBoundary = index
			}

			searchLeftBoundary = index

			// Move leftwards as long as characters match in the suffix.
			for searchRightBoundary >= 0 && pattern[searchRightBoundary] ==
				pattern[searchRightBoundary+patternLength-1-searchLeftBoundary] {
				searchRightBoundary--
			}

			// Length of the matching suffix at this position.
			suffixLengths[index] = searchLeftBoundary - searchRightBoundary
		}
	}

	// Return the array of suffix lengths.
	return suffixLengths
}

// preprocessGoodSuffixTable builds the good suffix shift table for the pattern.
// This table stores the amount by which the pattern should be shifted
// when a mismatch occurs after some matching suffix in the pattern.
// The shifts are computed using the suffix lengths and help optimize pattern alignment.
func preprocessGoodSuffixTable(pattern []rune) []int {
	var patternLength int = len(pattern)
	var goodSuffixShifts []int = make([]int, patternLength)

	var suffixLengths []int = preprocessSuffixes(pattern)

	// Initialize all shift values to pattern length (worst case).
	for index := range goodSuffixShifts {
		goodSuffixShifts[index] = patternLength
	}

	// Handle case where suffix matches prefix of the pattern.
	for index := patternLength - 1; index >= 0; index-- {
		if suffixLengths[index] == index+1 {
			for suffixPosition := 0; suffixPosition < patternLength-1-index; suffixPosition++ {
				// Only update shifts that are still at worst case.
				if goodSuffixShifts[suffixPosition] == patternLength {
					// Shift to align prefix with match suffix.
					goodSuffixShifts[suffixPosition] = patternLength - 1 - index
				}
			}
		}
	}

	// Handle other suffixes by assigning shifts based on suffix lengths.
	for index := 0; index <= patternLength-2; index++ {
		goodSuffixShifts[patternLength-1-suffixLengths[index]] = patternLength - 1 - index
	}

	// Return the completed good suffix shift table.
	return goodSuffixShifts
}

// BoyerMooreSearch performs the Boyer-Moore string search algorithm.
// It searches for all occurrences of the pattern in the given text
// and returns a slice of starting indices where the pattern is found.
// Utilizes both bad character and good suffix heuristics for efficient searching.
func BoyerMooreSearch(text string, pattern string) []int {
	var indices []int

	// Convert strings to rune slices to correctly handle Unicode.
	var textRunes []rune = []rune(text)
	var patternRunes []rune = []rune(pattern)

	var textLength int = len(textRunes)
	var patternLength int = len(patternRunes)

	// Return empty if pattern is empty or longer than the text.
	if patternLength == 0 || textLength < patternLength {
		return indices
	}

	// Preprocess tables used for efficient skipping.
	var badCharacterTable map[rune]int = preprocessBadCharacterTable(patternRunes)
	var goodSuffixShiftTable []int = preprocessGoodSuffixTable(patternRunes)

	var currentTextAlignment int = 0
	var patternIndex int = 0

	var mismatchedTextCharacter rune = ' '
	var lastKnownOccurrence int = 0

	var badCharacterShift int = 0
	var goodSuffixShift int = 0

	// Loop while pattern can still fit the remaining text.
	for currentTextAlignment <= textLength-patternLength {
		patternIndex = patternLength - 1

		// Compare pattern with text from end of pattern.
		for patternIndex >= 0 && patternRunes[patternIndex] == textRunes[currentTextAlignment+patternIndex] {
			patternIndex--
		}

		// Full match found.
		if patternIndex < 0 {
			indices = append(indices, currentTextAlignment)
			currentTextAlignment += goodSuffixShiftTable[0]
		} else {
			// Mismatch found, use heuristics to determine shift.
			mismatchedTextCharacter = textRunes[currentTextAlignment+patternIndex]
			lastKnownOccurrence = badCharacterTable[mismatchedTextCharacter]
			badCharacterShift = patternIndex - lastKnownOccurrence

			// Ensure at least one shift forward.
			if badCharacterShift < 1 {
				badCharacterShift = 1
			}

			// Calculate shift using good suffix rule.
			goodSuffixShift = goodSuffixShiftTable[patternIndex]

			// Shift by the maximum of both heuristics to maximize skipping.
			currentTextAlignment += maximum(badCharacterShift, goodSuffixShift)
		}
	}

	// Return all found indices.
	return indices
}
