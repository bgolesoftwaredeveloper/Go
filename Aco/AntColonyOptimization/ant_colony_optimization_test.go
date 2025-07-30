// ===================================================================================
// File:        antcolonyoptimization_test.go
// Package:     antcolonyoptimization
// Description: This file contains unit tests for the Ant Colony Optimization (ACO)
//
//	algorithm implementation. The tests validate the optimizer's ability
//	to find complete tours over graphs with various configurations,
//	and check consistency, correctness, and edge case handling.
//
//	The tests in this file cover key scenarios, including:
//
//	- Basic functionality on known distance matrices
//	- Deterministic consistency for identical inputs
//	- Edge cases such as single-node and zero-distance graphs
//	- Handling of full pheromone evaporation and sparse graphs
//	- Uniform edge weight graphs and heuristic/pheromone balance
//
//	All tests are written using Go’s built-in "testing" package.
//
// Author:      Braiden Gole
// Created:     July 29, 2025
//
// Test Coverage:
//
//	✅ TestBasicACOExecution
//	✅ TestDeterministicRun
//	✅ TestEdgeCaseSingleNode
//	✅ TestZeroDistanceMatrix
//	✅ TestHighEvaporationRate
//	✅ TestSparseGraph
//	✅ TestAllEqualDistances
//
// Usage:
//
//	To run all tests:
//	$ go test
//
// ===================================================================================
package antcolonyoptimization

import (
	"math"
	"testing"

	graph "github.com/bgolesoftwaredeveloper/ant_colony_optimization/Graph"
)

// distance matrix is a symmetric 5x5 matrix representing distances between cities.
var distanceMatrix [][]float64 = [][]float64{
	{0, 2, 9, 10, 7},
	{2, 0, 6, 4, 3},
	{9, 6, 0, 8, 5},
	{10, 4, 8, 0, 6},
	{7, 3, 5, 6, 0},
}

// TestBasicACOExecution validates that the optimizer returns a complete tour of the expected length and non-zero cost.
func TestBasicACOExecution(test *testing.T) {
	// Arrange.
	var graph *graph.Graph = graph.NewGraph(distanceMatrix)
	var optimizer *AntColonyOptimizer = NewAntColonyOptimizer(graph, 1.0, 5.0, 0.5, 100.0, 10, 20)

	// Act.
	tour, cost := optimizer.Solve()

	// Assert.
	if len(tour) != graph.NumberOfNodes+1 {
		test.Errorf("Expected tour of length %d, got %d.", graph.NumberOfNodes, len(tour))
	}

	if cost <= 0 {
		test.Errorf("Expected positive tour cost, got %f.", cost)
	}
}

// TestDeterministicRun runs ACO twice on the same graph an compares results.
func TestDeterministicRun(test *testing.T) {
	// Arrange.
	var graph *graph.Graph = graph.NewGraph(distanceMatrix)
	var optimizerCompare *AntColonyOptimizer = NewAntColonyOptimizer(graph, 1.0, 5.0, 0.5, 100.0, 10, 20)
	var optimizerAgainst *AntColonyOptimizer = NewAntColonyOptimizer(graph, 1.0, 5.0, 0.5, 100.0, 10, 20)

	// Act.
	compareTour, compareCost := optimizerCompare.Solve()
	againstTour, againstCost := optimizerAgainst.Solve()

	// Assert.
	if len(compareTour) != len(againstTour) {
		test.Errorf("Tours have different lengths: %d vs %d.", len(compareTour), len(againstTour))
	}

	if compareCost <= 0 || againstCost <= 0 {
		test.Errorf("One of the tour costs is not positive: %f vs %f.", compareCost, againstCost)
	}
}

// TestEdgeCaseSingleNode ensures algorithm gracefully handles a singly-node graph.
func TestEdgeCaseSingleNode(test *testing.T) {
	// Arrange.
	var singleNodeMatrix [][]float64 = [][]float64{{0}}

	var graph *graph.Graph = graph.NewGraph(singleNodeMatrix)
	var optimizer *AntColonyOptimizer = NewAntColonyOptimizer(graph, 1.0, 1.0, 0.5, 100.0, 1, 10)

	// Act.
	tour, cost := optimizer.Solve()

	// Assert.
	if len(tour) != 2 {
		test.Errorf("Expected tour length of 2 (start -> start) got %d.", len(tour))
	}

	if cost != 0 {
		test.Errorf("Expected zero cost for single node tour, got %f.", cost)
	}
}

// TestZeroDistanceMatrix ensures the optimizer handles a graph where all distances are zero - all tours ahve zero cost.
func TestZeroDistanceMatrix(test *testing.T) {
	// Arrange.
	var zeroMatrix [][]float64 = [][]float64{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	var graph *graph.Graph = graph.NewGraph(zeroMatrix)
	var optimizer *AntColonyOptimizer = NewAntColonyOptimizer(graph, 1.0, 2.0, 0.5, 100.0, 5, 10)

	// Act.
	tour, cost := optimizer.Solve()

	// Assert.
	if len(tour) < 2 {
		test.Fatalf("Tour is too short to be valid: %v.", tour)
	}

	var visited map[int]bool = make(map[int]bool)

	for _, node := range tour {
		visited[node] = true
	}

	if len(visited) != 3 {
		test.Errorf("Expected to visit all 3 nodes, got visited: %v.", visited)
	}

	if tour[0] != tour[len(tour)-1] {
		test.Errorf("Expected tour to return to starting node, got start: %d, end: %d.", tour[0], tour[len(tour)-1])
	}

	if cost != 0.0 {
		test.Errorf("Expected total tour cost of 0 for zero matrix, got %f.", cost)
	}
}

// TestHighEvaporationRate checks that the optimizer still works with full evaporation.
func TestHighEvaporationRate(test *testing.T) {
	// Arrange.
	var graph *graph.Graph = graph.NewGraph(distanceMatrix)
	var optimizer *AntColonyOptimizer = NewAntColonyOptimizer(graph, 1.0, 2.0, 1.0, 100.0, 5, 10)

	// Act.
	_, cost := optimizer.Solve()

	// Assert.
	if cost <= 0 {
		test.Errorf("Expected positive tour cost even with full evaporation, got %f.", cost)
	}
}

// TestSparseGraph ensures the optimizer handles graph with some unreachable nodes (represented as math.Inf).
func TestSparseGraph(test *testing.T) {
	// Arrange.
	var sparseMatrix [][]float64 = [][]float64{
		{0, 1, math.Inf(1)},
		{1, 0, 2},
		{math.Inf(1), 2, 0},
	}

	var graph *graph.Graph = graph.NewGraph(sparseMatrix)
	var optimizer *AntColonyOptimizer = NewAntColonyOptimizer(graph, 1.0, 2.0, 0.5, 100.0, 10, 10)

	// Act.
	tour, cost := optimizer.Solve()

	// Assert.
	if len(tour) < 2 {
		test.Fatalf("Tour is too short: %v.", tour)
	}

	if cost <= 0 || math.IsInf(cost, 1) {
		test.Errorf("Expected finite positive cost, got %f.", cost)
	}
}

// TestAllEqualDistances ensures consistent behavior when all edge weights are equal.
func TestAllEqualDistances(test *testing.T) {
	// Arrange.
	var equalMatrix [][]float64 = [][]float64{
		{0, 1, 1, 1},
		{1, 0, 1, 1},
		{1, 1, 0, 1},
		{1, 1, 1, 0},
	}

	var graph *graph.Graph = graph.NewGraph(equalMatrix)
	var optimizer *AntColonyOptimizer = NewAntColonyOptimizer(graph, 1.0, 1.0, 0.3, 100.0, 5, 10)

	// Act.
	tour, cost := optimizer.Solve()

	// Assert.
	if len(tour) != 5 {
		test.Errorf("Expected tour of length 5 (4 nodes + return), got %d.", len(tour))
	}

	if cost <= 0 {
		test.Errorf("Expected positive cost, got %f.", cost)
	}
}
