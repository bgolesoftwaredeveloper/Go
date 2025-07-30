// ===================================================================================
// File:        main.go
// Package:     main
// Description: Entry point for the Ant Colony Optimization (ACO) example program.
//
//	This program sets up a simple TSP-like problem with a predefined
//	distance matrix representing a small network of cities. It creates
//	a graph from the distance matrix, configures the Ant Colony Optimizer
//	with parameters, runs the optimization to find the best tour,
//	and prints the resulting best path and cost.
//
// Author:      Braiden Gole
// Created:     July 29, 2025
//
// ===================================================================================
package main

import (
	"fmt"

	optimization "github.com/bgolesoftwaredeveloper/ant_colony_optimization/AntColonyOptimization"
	graph "github.com/bgolesoftwaredeveloper/ant_colony_optimization/Graph"
)

func main() {
	// Define the distance matrix representing the weighted edges between nodes (cities).
	var distanceMatrix [][]float64 = [][]float64{
		{0, 2, 9, 10, 7},
		{2, 0, 6, 4, 3},
		{9, 6, 0, 8, 5},
		{10, 4, 8, 0, 1},
		{7, 3, 5, 1, 0},
	}

	// Create a new graph instance with the distance matrix.
	var cityMap *graph.Graph = graph.NewGraph(distanceMatrix)

	// Initialize the Ant Colony Optimizer with problem graph and parameters:
	// alpha = 1.0 (pheromone influence),
	// beta = 5.0 (heuristic influence),
	// evaporation rate = 0.5,
	// deposit factor = 100.0,
	// number of ants = 10,
	// number of epochs = 100.
	var optimizer *optimization.AntColonyOptimizer = optimization.NewAntColonyOptimizer(
		cityMap,
		1.0,
		5.0,
		0.5,
		100.0,
		10,
		100,
	)

	// Run the ACO solver to find the best tour and its cost.
	bestPath, bestCost := optimizer.Solve()

	// Ouput the best tour found ant its total cost.
	fmt.Println("Best tour found:", bestPath)
	fmt.Println("Tour cost:", bestCost)
}
