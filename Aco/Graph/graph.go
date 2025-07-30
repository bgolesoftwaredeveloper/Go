// ===================================================================================
// File:        graph.go
// Package:     graph
// Description: This package provides the Graph type used to represent a problem graph
//
//	for the Ant Colony Optimization (ACO) algorithm.
//
//	The Graph structure stores the number of nodes and a distance matrix,
//	where each entry represents the distance between two nodes.
//
//	Key functionalities include:
//	- Creating a new Graph from a given distance matrix
//	- Querying the distance between two nodes
//	- Calculating Euclidean distance between two points (utility function)
//
// Author:      Braiden Gole
// Created:     July 29, 2025
//
// ===================================================================================
package graph

// Graph represents a weighted graph with a distance matrix.
//
// NumberOfNodes    - the total count of nodes in the graph
// DistanceMatrix   - a 2D slice storing distances between nodes;
//                    DistanceMatrix[i][j] gives the distance from node i to j
type Graph struct {
	NumberOfNodes  int
	DistanceMatrix [][]float64
}

// NewGraph constructs a new Graph instance using the provided distance matrix.
//
// Parameters:
//   distanceMatrix - a 2D slice representing distances between nodes;
//                    must be square (NxN) where N is number of nodes.
//
// Returns:
//   Pointer to the newly created Graph.
func NewGraph(distanceMatrix [][]float64) *Graph {
	return &Graph{
		NumberOfNodes:  len(distanceMatrix),
		DistanceMatrix: distanceMatrix,
	}
}

// DistanceBetween returns the distance between the source and destination nodes.
//
// Parameters:
//   source      - the index of the source node
//   destination - the index of the destination node
//
// Returns:
//   The distance as a float64 value.
func (graph *Graph) DistanceBetween(source int, destination int) float64 {
	return graph.DistanceMatrix[source][destination]
}

// EclideanDistance calculates the Euclidean distance between two points in 2D space.
//
// Parameters:
//   xAxis       - x-coordinate of the first point
//   yAxis       - y-coordinate of the first point
//   otherXAxis  - x-coordinate of the second point
//   otherYAxis  - y-coordinate of the second point
//
// Returns:
//   The Euclidean distance as a float64 value.
/*
func EclideanDistance(xAxis float64, yAxis float64, otherXAxis float64, otherYAxis float64) float64 {
	return math.Hypot(otherXAxis-xAxis, otherYAxis-xAxis)
}
*/
