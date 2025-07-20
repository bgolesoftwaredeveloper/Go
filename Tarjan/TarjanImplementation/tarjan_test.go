// ===================================================================================
// File:        tarjan_test.go
// Package:     tarjanimplementation
// Description: This file contains unit tests for the Tarjan algorithm implementation.
//
//	Tarjan's algorithm is used to identify all strongly connected components (SCCs)
//	in a directed graph using depth-first search and low-link value comparisons.
//
//	The tests in this file verify the correctness and robustness of the SCC
//	detection logic across various graph configurations, including:
//
//	- Disconnected graphs
//	- Cyclic and acyclic graphs
//	- Self-loops and single-node components
//	- Complex intertwined components
//
//	All tests are written using Go’s built-in "testing" package.
//
// Author:      Braiden Gole
// Created:     July 20, 2025
//
// Test Coverage:
//
//	✅ TestSinglyComponent
//	✅ TestDisconnectedVertices
//	✅ TestMultipleComponents
//	✅ TestGraphWithSelfLoop
//	✅ TestEmptyGraph
//	✅ TestLinearGraphNoCycles
//	✅ TestComponentWithBackEdge
//
// Usage:
//
//	To run all tests:
//	$ go test
//
// ===================================================================================
package tarjanimplementation

import (
	"reflect"
	"sort"
	"testing"
)

// ===============================
// Utility: SCC Equality Assertion
// ===============================
func assertComponentEqual(test *testing.T, actual [][]int, expected [][]int) {
	// Normalize inner slices and overall slice ordering.
	var normalize = func(components [][]int) [][]int {
		for index := range components {
			sort.Ints(components[index])
		}

		sort.Slice(components, func(rowIndex, columnIndex int) bool {
			return components[rowIndex][0] < components[columnIndex][0]
		})

		return components
	}

	// Normalize both expected and actual before comparing.
	actual = normalize(actual)
	expected = normalize(expected)

	// Assert.
	if !reflect.DeepEqual(actual, expected) {
		test.Errorf("Strongly connected components mismatch.\nExpected: %v\nGot:\t %v.", expected, actual)
	}
}

// TestSinglyComponent verifies that a single cycle is detected as one component.
func TestSinglyComponent(test *testing.T) {
	// Arrange.
	var graph map[int][]int = map[int][]int{
		1: {2},
		2: {3},
		3: {1},
	}

	var expected [][]int = [][]int{{1, 2, 3}}

	// Act.
	var finder *TarjanStronglyConnectedComponent = NewTarjanStronglyConnectedComponent(graph)
	var result [][]int = finder.FindStronglyConnectedComponents()

	// Assert.
	assertComponentEqual(test, result, expected)
}

// TestDisconnetedVertices check that isolated nodes are detected as individual SCC.
func TestDisconnectedVertices(test *testing.T) {
	// Arrange.
	var graph map[int][]int = map[int][]int{
		1: {},
		2: {},
		3: {},
	}

	var expected [][]int = [][]int{{1}, {2}, {3}}

	// Act.
	var finder *TarjanStronglyConnectedComponent = NewTarjanStronglyConnectedComponent(graph)
	var result [][]int = finder.FindStronglyConnectedComponents()

	// Assert.
	assertComponentEqual(test, result, expected)
}

// TestMultipleComponents verifies SCC detection across multiple interconnected cycles.
func TestMultipleComponents(test *testing.T) {
	// Arrange.
	var graph map[int][]int = map[int][]int{
		0: {1},
		1: {2},
		2: {0},
		3: {4},
		4: {5},
		5: {3},
		6: {},
	}

	var expected [][]int = [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6},
	}

	// Act.
	var finder *TarjanStronglyConnectedComponent = NewTarjanStronglyConnectedComponent(graph)
	var result [][]int = finder.FindStronglyConnectedComponents()

	// Assert.
	assertComponentEqual(test, result, expected)
}

// TestGraphWithSelfLoop confirms that self-loops form their own SCC.
func TestGraphWithSelfLoop(test *testing.T) {
	// Arrange.
	var graph map[int][]int = map[int][]int{
		1: {1},
		2: {3},
		3: {2},
	}

	var expected [][]int = [][]int{
		{1},
		{2, 3},
	}

	// Act.
	var finder *TarjanStronglyConnectedComponent = NewTarjanStronglyConnectedComponent(graph)
	var result [][]int = finder.FindStronglyConnectedComponents()

	// Assert.
	assertComponentEqual(test, result, expected)
}

// TestEmptyGraph ensures no components are retunr from an empty graph.
func TestEmptyGraph(test *testing.T) {
	// Arrange.
	var graph map[int][]int = map[int][]int{}
	var expected [][]int = [][]int{}

	// Act.
	var finder *TarjanStronglyConnectedComponent = NewTarjanStronglyConnectedComponent(graph)
	var result [][]int = finder.FindStronglyConnectedComponents()

	// Assert.
	assertComponentEqual(test, result, expected)
}

// TestLinearGraphWithNoCycles checks that a linear DAG results in one-node SCCs.
func TestLinearGraphNoCycles(test *testing.T) {
	// Arrange.
	var graph map[int][]int = map[int][]int{
		1: {2},
		2: {3},
		3: {},
	}

	var expected [][]int = [][]int{{3}, {2}, {1}}

	// Act.
	var finder *TarjanStronglyConnectedComponent = NewTarjanStronglyConnectedComponent(graph)
	var result [][]int = finder.FindStronglyConnectedComponents()

	// Assert.
	assertComponentEqual(test, result, expected)
}

// TestComponentsWithBackEdge ensures that a back edge creates a unified SCC.
func TestComponentsWithBackEdge(test *testing.T) {
	// Arrange.
	var graph map[int][]int = map[int][]int{
		1: {2},
		2: {3},
		3: {4},
		// Back edge forming a cycle: 2 → 3 → 4 → 2
		4: {2},
		5: {},
	}

	var expected [][]int = [][]int{
		{2, 3, 4},
		{1},
		{5},
	}

	// Act.
	var finder *TarjanStronglyConnectedComponent = NewTarjanStronglyConnectedComponent(graph)
	var result [][]int = finder.FindStronglyConnectedComponents()

	// Assert.
	assertComponentEqual(test, result, expected)
}
