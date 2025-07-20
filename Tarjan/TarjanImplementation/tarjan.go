// ===================================================================================
// File:        tarjan.go
// Package:     tarjan
// Description: This package implements Tarjan's algorithm in Go to identify strongly
//
//	connected components (SCCs) in a directed graph. The algorithm uses
//	depth-first search (DFS) and low-link values to efficiently identify SCCs in
//	O(V + E) time complexity.
//
//	Features implemented in this package:
//	- TarjanSCCFinder struct to encapsulate internal state
//	- Recursive depth-first search and low-link comparison
//	- Stack tracking to manage component membership
//	- Returns a slice of SCCs, where each SCC is a slice of vertex IDs
//
// Author:      Braiden Gole
// Created:     July 20, 2025
//
// ===================================================================================
package tarjanimplementation

// Tarjan strongly connected component holds the internal state used during SCC detection.
type TarjanStronglyConnectedComponent struct {
	graph                       map[int][]int
	index                       int
	nodeIndex                   map[int]int
	lowLinkValue                map[int]int
	onStack                     map[int]bool
	stack                       []int
	stronglyConnectedComponents [][]int
}

// NewTarjanStronglyConnectedComponent initializes a new TarjanStronglyConnectedComponent with the provided graph.
func NewTarjanStronglyConnectedComponent(graph map[int][]int) *TarjanStronglyConnectedComponent {
	return &TarjanStronglyConnectedComponent{
		graph:                       graph,
		index:                       0,
		nodeIndex:                   make(map[int]int),
		lowLinkValue:                make(map[int]int),
		onStack:                     make(map[int]bool),
		stack:                       []int{},
		stronglyConnectedComponents: [][]int{},
	}
}

// strongConnect is a recursive helper that performs the DFS and identifies strongly connected components based on index and
// low-link comparisons.
func (tarjan *TarjanStronglyConnectedComponent) strongConnect(vertex int) {
	// Assign discovery index and low-link value to the current vertex.
	tarjan.nodeIndex[vertex] = tarjan.index
	tarjan.lowLinkValue[vertex] = tarjan.index
	tarjan.index++

	// Push the vertex onto the stack and mark it as "on stack."
	tarjan.stack = append(tarjan.stack, vertex)
	tarjan.onStack[vertex] = true

	// Explore all adjacent vertices.
	for _, neighbor := range tarjan.graph[vertex] {
		// If the neighbor has not been visited, recurse on it.
		if _, visited := tarjan.nodeIndex[neighbor]; !visited {
			tarjan.strongConnect(neighbor)

			// Update the low-link value based on the recursive result.
			if tarjan.lowLinkValue[neighbor] < tarjan.lowLinkValue[vertex] {
				tarjan.lowLinkValue[vertex] = tarjan.lowLinkValue[neighbor]
			}
		} else if tarjan.onStack[neighbor] {
			// If the neighbor is on the stack, it is part of the current component.
			// Update low-link based on the discovery index of the neighbor.
			if tarjan.nodeIndex[neighbor] < tarjan.lowLinkValue[vertex] {
				tarjan.lowLinkValue[vertex] = tarjan.nodeIndex[neighbor]
			}
		}
	}

	// If the current vertex is a root of an SCC.
	if tarjan.lowLinkValue[vertex] == tarjan.nodeIndex[vertex] {
		var component []int

		// Pop vertices from the stack to form the strongly connected components.
		for {
			var popped int = tarjan.stack[len(tarjan.stack)-1]

			tarjan.stack = tarjan.stack[:len(tarjan.stack)-1]
			tarjan.onStack[popped] = false

			component = append(component, popped)

			// Stop when the current root vertex is reached.
			if popped == vertex {
				break
			}
		}

		// Append the identified component to the result list.
		tarjan.stronglyConnectedComponents = append(tarjan.stronglyConnectedComponents, component)
	}
}

// FindStronglyConnectedComponents executes Tarjan's algorithm and returns a slice of strongly connected components.
// Each strongly connected component is represented as a slice of integers.
func (tarjan *TarjanStronglyConnectedComponent) FindStronglyConnectedComponents() [][]int {
	// Visit all vertices in the graph. Start DFS if the vertex has not been visited yet.
	for vertex := range tarjan.graph {
		if _, visited := tarjan.nodeIndex[vertex]; !visited {
			tarjan.strongConnect(vertex)
		}
	}

	// Return the complete list of strongly connected components.
	return tarjan.stronglyConnectedComponents
}
