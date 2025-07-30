// ===================================================================================
// File:        antcolonyoptimization.go
// Package:     antcolonyoptimization
// Description: This package implements the AntColonyOptimizer type, which encapsulates
//
//	the core logic of the Ant Colony Optimization (ACO) metaheuristic.
//
//	The AntColonyOptimizer manages the problem graph, pheromone levels,
//	and algorithm parameters (alpha, beta, evaporation rate, deposit factor).
//	It simulates multiple ants constructing solutions over many epochs,
//	updating pheromones to guide future path selections toward optimal tours.
//
//	Key features:
//	- Initialization with problem graph and parameters
//	- Running the optimization to find a near-optimal tour
//	- Pheromone evaporation and deposition to balance exploration/exploitation
//
// Author:      Braiden Gole
// Created:     July 29, 2025
//
// ===================================================================================
package antcolonyoptimization

import (
	"math"
	"math/rand"

	ant "github.com/bgolesoftwaredeveloper/ant_colony_optimization/Ant"
	graph "github.com/bgolesoftwaredeveloper/ant_colony_optimization/Graph"
	pheromone "github.com/bgolesoftwaredeveloper/ant_colony_optimization/Pheromone"
)

// AntColonyOptimizer encapsulates the parameters and state needed to run the
// Ant Colony Optimization algorithm.
//
// ProblemGraph    - the graph representing the problem to be solved
// PheromoneLevels - matrix tracking pheromone intensities on graph edges
// Alpha           - influence weight of pheromone strength on path selection
// Beta            - influence weight of heuristic visibility (inverse distance) on path selection
// EvaporateRate   - rate at which pheromone evaporates each epoch (decay factor)
// DepositFactor   - scaling factor for pheromone deposited by ants after tours
// NumberOfAnts    - number of ants constructing tours each epoch
// NumberOfEpochs  - number of iterations to run the optimization process
type AntColonyOptimizer struct {
	ProblemGraph    *graph.Graph
	PheromoneLevels *pheromone.PheromoneMatrix
	Alpha           float64
	Beta            float64
	EvaporateRate   float64
	DepositFactor   float64
	NumberOfAnts    int
	NumberOfEpochs  int
}

// NewAntColonyOptimizer initializes and returns a new AntColonyOptimizer instance.
//
// Parameters:
//
//	graph          - the problem graph to solve
//	alpha          - weight of pheromone influence on path selection
//	beta           - weight of heuristic (visibility) influence on path selection
//	evaporationRate- pheromone evaporation rate per epoch (0.0 to 1.0)
//	depositFactor  - scaling factor for pheromone deposit amount
//	antCount       - number of ants per epoch
//	epochCount     - number of epochs (iterations) to run
//
// Returns:
//
//	Pointer to a fully initialized AntColonyOptimizer.
func NewAntColonyOptimizer(graph *graph.Graph, alpha, beta, evaporationRate, depositFactor float64, antCount, epochCount int) *AntColonyOptimizer {
	var pheromones *pheromone.PheromoneMatrix = pheromone.NewPheromoneMatrix(graph.NumberOfNodes, 1.0)

	return &AntColonyOptimizer{
		ProblemGraph:    graph,
		PheromoneLevels: pheromones,
		Alpha:           alpha,
		Beta:            beta,
		EvaporateRate:   evaporationRate,
		DepositFactor:   depositFactor,
		NumberOfAnts:    antCount,
		NumberOfEpochs:  epochCount,
	}
}

// Solve executes the ACO algorithm over the configured number of epochs,
// simulating ants constructing tours, updating pheromones, and tracking
// the best tour found.
//
// Returns:
//
//	bestTour     - slice of node indices representing the best tour found
//	bestTourCost - total cost (distance) of the best tour
func (antColonyOptimizer *AntColonyOptimizer) Solve() ([]int, float64) {
	var bestTour []int = []int{}
	var bestTourCost float64 = math.MaxFloat64

	var ants []*ant.Ant
	var currentAnt *ant.Ant

	for epoch := 0; epoch < antColonyOptimizer.NumberOfEpochs; epoch++ {
		ants = make([]*ant.Ant, antColonyOptimizer.NumberOfAnts)

		for index := 0; index < antColonyOptimizer.NumberOfAnts; index++ {
			currentAnt = ant.NewAnt(antColonyOptimizer.ProblemGraph, antColonyOptimizer.PheromoneLevels,
				antColonyOptimizer.Alpha, antColonyOptimizer.Beta)

			// Construct a tour starting from a random node.
			currentAnt.ConstructTour(rand.Intn(antColonyOptimizer.ProblemGraph.NumberOfNodes))

			ants[index] = currentAnt

			// Update best solution found so far.
			if currentAnt.TotalCost < bestTourCost {
				bestTourCost = currentAnt.TotalCost
				bestTour = append([]int(nil), currentAnt.PathTaken...)
			}
		}

		// Evaporate pheromones to simulate natural decay.
		antColonyOptimizer.PheromoneLevels.Evaporate(antColonyOptimizer.EvaporateRate)

		// Deposit pheromones based on ant tours, reinforcing shorter paths.
		for _, insect := range ants {
			antColonyOptimizer.PheromoneLevels.DepositPheromones(insect.PathTaken, antColonyOptimizer.DepositFactor/insect.TotalCost)
		}
	}

	return bestTour, bestTourCost
}
