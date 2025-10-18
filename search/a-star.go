package search

import (
	"sync"
)

// Node do grafo
type Node struct {
	ID        int
	Neighbors map[int]int
}

type GraphInput struct {
	Nodes map[int]*Node
}

// Dijkstra: Concorrencia com controle de threads
func ConcurrentDijkstra(graph *GraphInput, start, goal int, maxThreads int) (int, bool) {
	visited := sync.Map{}
	costSoFar := sync.Map{}

	costSoFar.Store(start, 0)

	sem := make(chan struct{}, maxThreads)
	var wg sync.WaitGroup

	result := make(chan int, 1)
	found := false

	var explore func(currentID int, currentCost int)

	explore = func(currentID int, currentCost int) {
		defer wg.Done()

		// limita qtde. de threads concorrentes
		sem <- struct{}{}
		defer func() { <-sem }()

		if currentID == goal {
			if !found {
				found = true
				result <- currentCost
			}
			return
		}

		visited.Store(currentID, true)

		node := graph.Nodes[currentID]

		for neighbor, weight := range node.Neighbors {
			if _, ok := visited.Load(neighbor); ok {
				continue
			}

			newCost := currentCost + weight
			prevCostInterface, exists := costSoFar.Load(neighbor)
			prevCost := 0
			if exists {
				prevCost = prevCostInterface.(int)
			}

			if !exists || newCost < prevCost {
				costSoFar.Store(neighbor, newCost)
				wg.Add(1)
				go explore(neighbor, newCost)
			}
		}
	}

	wg.Add(1)
	go explore(start, 0)
	wg.Wait()
	close(result)

	if len(result) > 0 {
		return <-result, true
	}
	return 0, false
}
