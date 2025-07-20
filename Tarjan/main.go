// ===================================================================================
// File:        main.go
// Package:     main
// Description: This file demonstrates usage of the Tarjan algorithm implementation
//
//	by constructing a sample directed graph and identifying all strongly
//	connected components using the tarjan package.
//
// Author:      Braiden Gole
// Created:     July 20, 2025
//
// Example Usage:
//
//	go run main.go
//
// ===================================================================================
package main

import (
	"fmt"

	tarjan "github.com/bgolesoftwaredeveloper/tarjan/TarjanImplementation"
)

func main() {
	// =================================================================================
	// Step 1: Define the input graph as an adjacency list.
	//
	// The graph is a map[int][]int where each key represents a vertex,
	// and its value is a slice of all vertices it points to.
	//
	// Example graph:
	//
	//	    0 → 1
	//	    ↑   ↓
	//	    2 ← 3 → 4 → 5 → 3
	//
	// This graph has two strongly connected components:
	// - [0 1 2]
	// - [3 4 5]
	// =================================================================================
	var graph map[int][]int = map[int][]int{
		0: {1},
		1: {2},
		2: {0},
		3: {1, 4},
		4: {5},
		5: {3},
	}

	// =================================================================================
	// Step 2: Create a new TarjanStronglyConnectedComponent object using the graph.
	// =================================================================================
	var finder *tarjan.TarjanStronglyConnectedComponent = tarjan.NewTarjanStronglyConnectedComponent(graph)

	// =================================================================================
	// Step 3: Find all strongly connected components in the graph.
	// =================================================================================
	var stronglyConnectedComponents [][]int = finder.FindStronglyConnectedComponents()

	// =================================================================================
	// Step 4: Display the results.
	// Each component is a slice of vertex IDs.
	// =================================================================================
	fmt.Println("Strongly connected components:")

	for index, component := range stronglyConnectedComponents {
		fmt.Printf("\tComponent %d: %v\n", index+1, component)
	}
}
