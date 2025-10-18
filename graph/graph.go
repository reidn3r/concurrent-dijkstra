package graph

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	ID        int
	Neighbors map[int]int
}

type Graph struct {
	Nodes map[int]*Node
}

// Grafo aleatório
func GenerateGraph(n int, maxEdgesPerNode int, maxWeight int) *Graph {
	rand.Seed(time.Now().UnixNano())
	g := &Graph{Nodes: make(map[int]*Node)}

	// criar nós
	for i := 0; i < n; i++ {
		g.Nodes[i] = &Node{
			ID:        i,
			Neighbors: make(map[int]int),
		}
	}

	// criar arestas aleatórias
	for _, node := range g.Nodes {
		numEdges := rand.Intn(maxEdgesPerNode) + 1
		for j := 0; j < numEdges; j++ {
			neighborID := rand.Intn(n)
			if neighborID != node.ID {
				weight := rand.Intn(maxWeight) + 1
				node.Neighbors[neighborID] = weight
			}
		}
	}

	return g
}

// Printar o grafo
func (g *Graph) Print() {
	for id, node := range g.Nodes {
		fmt.Printf("Node %d: ", id)
		for n, w := range node.Neighbors {
			fmt.Printf("(%d,%d) ", n, w)
		}
		fmt.Println()
	}
}
