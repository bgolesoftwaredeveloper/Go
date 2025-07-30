// ===================================================================================
// File:        pheromone.go
// Package:     pheromone
// Description: This package implements the PheromoneMatrix type used to represent and
//
//	manage pheromone levels in an Ant Colony Optimization (ACO) algorithm.
//
//	The PheromoneMatrix maintains a 2D matrix of float64 values representing pheromone
//	intensities on edges between nodes in a graph. It supports key operations:
//
//	- Initialization with a given size and initial pheromone value
//	- Evaporation of pheromone levels by a specified rate to simulate decay over time
//	- Depositing pheromones along a given path, increasing pheromone levels on edges
//
//	This structure is essential for controlling the probabilistic path selection of ants
//	in the ACO metaheuristic by dynamically adjusting edge desirability.
//
// Author:      Braiden Gole
// Created:     July 29, 2025
//
// ===================================================================================
package pheromone

// PheromoneMatrix represents a 2D matrix of pheromone levels for edges between nodes
// in a graph, used in Ant Colony Optimization (ACO) algorithms.
//
// Each entry Values[i][j] holds the pheromone intensity on the edge from node i to node j.
// The matrix is symmetric as pheromones are deposited bidirectionally.
//
// This structure supports initialization with a uniform pheromone value, evaporation to
// simulate pheromone decay, and pheromone deposition along ant traversal paths to guide
// future ant decisions.
//
// By dynamically updating pheromone levels, the PheromoneMatrix helps balance exploration
// and exploitation in finding optimized paths on the problem graph.
type PheromoneMatrix struct {
	Values [][]float64
}

// NewPheromoneMatrix creates and initializes a new PheromoneMatrix with the specified
// number of nodes. Each edge's pheromone level is set to the provided initialValue.
//
// Parameters:
//   nodeCount    - the number of nodes in the graph (matrix size)
//   initialValue - the initial pheromone level for all edges
//
// Returns:
//   Pointer to the newly created PheromoneMatrix.
func NewPheromoneMatrix(nodeCount int, initialValue float64) *PheromoneMatrix {
	var matrix [][]float64 = make([][]float64, nodeCount)

	for row := range matrix {
		matrix[row] = make([]float64, nodeCount)

		for column := range matrix[row] {
			matrix[row][column] = initialValue
		}
	}

	return &PheromoneMatrix{Values: matrix}
}

// Evaporate reduces the pheromone levels on all edges by the given evaporation rate.
// This simulates pheromone decay over time, encouraging exploration.
//
// Parameters:
//   evaporationRate - the fraction of pheromone to evaporate (e.g., 0.1 reduces pheromone by 10%)
func (matrix *PheromoneMatrix) Evaporate(evaporationRate float64) {
	for row := range matrix.Values {
		for column := range matrix.Values[row] {
			matrix.Values[row][column] *= (1.0 - evaporationRate)
		}
	}
}

// DepositPheromones adds pheromone amounts along the edges defined by the given path.
// Both directions of each edge are incremented to maintain symmetry.
//
// Parameters:
//   path          - slice of node indices representing the path taken by an ant
//   depositAmount - the amount of pheromone to deposit on each edge along the path
func (matrix *PheromoneMatrix) DepositPheromones(path []int, depositAmount float64) {
	var from int
	var to int

	for index := 0; index < len(path)-1; index++ {
		from = path[index]
		to = path[index+1]

		matrix.Values[from][to] += depositAmount
		matrix.Values[to][from] += depositAmount
	}
}
