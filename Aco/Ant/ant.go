// ===================================================================================
// File:        ant.go
// Package:     ant
// Description: This package implements the Ant type used in the Ant Colony Optimization (ACO)
//
//	metaheuristic algorithm.
//
//	The Ant type simulates an individual ant that constructs a solution path
//	(tour) over a given graph based on pheromone intensity and heuristic visibility.
//
//	Key functionalities include:
//	- Tracking visited nodes and the path taken during a tour
//	- Selecting the next node to visit probabilistically using pheromone and distance info
//	- Constructing a complete tour starting from a root node and returning to it
//
//	This package works closely with the Graph package (problem graph representation)
//	and the Pheromone package (pheromone matrix managing edge desirability).
//
// Author:      Braiden Gole
// Created:     July 29, 2025
//
// ===================================================================================
package ant

import (
	"math"
	"math/rand"
	"time"

	graph "github.com/bgolesoftwaredeveloper/ant_colony_optimization/Graph"
	pheromone "github.com/bgolesoftwaredeveloper/ant_colony_optimization/Pheromone"
)

// Ant represents a single ant in the Ant Colony Optimization algorithm.
//
// It tracks visited nodes, the path taken, total cost of the tour, and
// references to the problem graph and pheromone matrix. The parameters alpha
// and beta control the influence of pheromone intensity and visibility
// (heuristic information) when selecting the next node.
type Ant struct {
	visitedNodes map[int]bool
	PathTaken    []int
	TotalCost    float64
	problemGraph *graph.Graph
	pheromones   *pheromone.PheromoneMatrix
	alpha        float64
	beta         float64
}

var randomNumberGenerator *rand.Rand

func init() {
	randomNumberGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// NewAnt creates and initializes a new Ant instance with the given problem graph,
// pheromone matrix, and parameters alpha and beta that weight pheromone and heuristic.
//
// Parameters:
//
//	graph      - the problem graph
//	pheromones - pheromone matrix controlling pheromone levels on edges
//	alpha      - influence of pheromone strength on path selection
//	beta       - influence of heuristic visibility on path selection
//
// Returns:
//
//	Pointer to the newly created Ant instance.
func NewAnt(graph *graph.Graph, pheromones *pheromone.PheromoneMatrix, alpha, beta float64) *Ant {
	return &Ant{
		visitedNodes: make(map[int]bool),
		PathTaken:    make([]int, 0, graph.NumberOfNodes),
		TotalCost:    0.0,
		problemGraph: graph,
		pheromones:   pheromones,
		alpha:        alpha,
		beta:         beta,
	}
}

// SelectNextNode chooses the next node for the ant to move to from the current node.
//
// It calculates the probability of moving to each unvisited neighbor based on pheromone
// levels raised to the power alpha and heuristic visibility raised to the power beta.
// Then, it performs roulette wheel selection to probabilistically select the next node.
//
// Parameters:
//
//	currentNode - the node where the ant currently is
//
// Returns:
//
//	The index of the selected next node, or -1 if no valid moves are available.
func (ant *Ant) SelectNextNode(currentNode int) int {
	var nodeCount int = ant.problemGraph.NumberOfNodes

	// Slice to hold move probabilities for each node.
	var probabilityList []float64 = make([]float64, nodeCount)

	var probabilitySum float64 = 0.0
	var pheromoneStrength float64 = 0.0
	var distance float64 = 0.0
	var visibility float64 = 0.0

	const EPSILON float64 = 1e-10

	for nextNode := 0; nextNode < nodeCount; nextNode++ {
		// Skip nodes already visited or the current node itself.
		if ant.visitedNodes[nextNode] || nextNode == currentNode {
			continue
		}

		pheromoneStrength = math.Pow(ant.pheromones.Values[currentNode][nextNode], ant.alpha)
		distance = ant.problemGraph.DistanceBetween(currentNode, nextNode)
		visibility = math.Pow(1.0/(distance+EPSILON), ant.beta)

		probabilityList[nextNode] = pheromoneStrength * visibility
		probabilitySum += probabilityList[nextNode]
	}

	// Safe check to avoid division by zero.
	if probabilitySum == 0 {
		return -1
	}

	// Normalize probabilities.
	for index := 0; index < nodeCount; index++ {
		probabilityList[index] /= probabilitySum
	}

	// Roulette wheel selection.
	var randomValue = randomNumberGenerator.Float64()
	var cumulativeProbability float64 = 0.0

	for index, probability := range probabilityList {
		cumulativeProbability += probability

		if randomValue <= cumulativeProbability {
			return index
		}
	}

	// Fallback: No valid node selected.
	return -1
}

// ConstructTour builds a complete tour for the ant starting from rootNode.
//
// The ant repeatedly selects the next node probabilistically until all nodes are visited,
// then returns to the root node to complete the cycle. It tracks the path taken and
// accumulates the total cost of the tour.
//
// Parameters:
//
//	rootNode - the starting node for the ant's tour
func (ant *Ant) ConstructTour(rootNode int) {
	// Reset states.
	ant.visitedNodes = make(map[int]bool)
	ant.PathTaken = ant.PathTaken[:0]
	ant.TotalCost = 0.0

	ant.PathTaken = append(ant.PathTaken, rootNode)
	ant.visitedNodes[rootNode] = true

	var currentNode int = rootNode
	var nextNode int = 0

	for len(ant.PathTaken) < ant.problemGraph.NumberOfNodes {
		nextNode = ant.SelectNextNode(currentNode)

		if nextNode == -1 {
			break
		}

		ant.PathTaken = append(ant.PathTaken, nextNode)
		ant.visitedNodes[nextNode] = true
		ant.TotalCost += ant.problemGraph.DistanceBetween(currentNode, nextNode)

		currentNode = nextNode
	}

	// Return to root node.
	ant.PathTaken = append(ant.PathTaken, rootNode)
	ant.TotalCost += ant.problemGraph.DistanceBetween(currentNode, rootNode)
}
