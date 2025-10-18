package main

import (
	"dijkstra/graph"
	"dijkstra/search"
	"fmt"
	"time"
)

func main() {
	numNodes := 100
	maxEdges := 5
	maxWeight := 10
	startNode := 0
	goalNode := numNodes - 1
	maxThreads := 15

	g := graph.GenerateGraph(numNodes, maxEdges, maxWeight)
	g.Print()

	graphInput := &search.GraphInput{
		Nodes: make(map[int]*search.Node),
	}
	for id, n := range g.Nodes {
		graphInput.Nodes[id] = &search.Node{
			ID:        n.ID,
			Neighbors: n.Neighbors,
		}
	}

	// Tempo de execução
	startTime := time.Now()
	cost, found := search.ConcurrentDijkstra(graphInput, startNode, goalNode, maxThreads)
	duration := time.Since(startTime)

	if found {
		fmt.Printf("Caminho encontrado -- Custo total: %d\n", cost)
	} else {
		fmt.Println("Caminho não encontrado.")
	}

	fmt.Printf("Tempo total de execução: %s\n", duration)
}
