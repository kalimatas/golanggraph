package main

import (
	_ "fmt"
	"github.com/kalimatas/graph"
	_ "github.com/kalimatas/structures"
)

func main() {
	theGraph := new(graph.Graph)

	theGraph.AddVertex(&graph.Vertex{Label: "A"}) // 0
	theGraph.AddVertex(&graph.Vertex{Label: "B"}) // 1
	theGraph.AddVertex(&graph.Vertex{Label: "C"}) // 2
	theGraph.AddVertex(&graph.Vertex{Label: "D"}) // 3
	theGraph.AddVertex(&graph.Vertex{Label: "E"}) // 4
	theGraph.AddVertex(&graph.Vertex{Label: "F"}) // 5

	theGraph.AddEdge(0, 1) // AB
	theGraph.AddEdge(1, 2) // BC
	theGraph.AddEdge(0, 3) // AD
	theGraph.AddEdge(3, 4) // DE
	theGraph.AddEdge(4, 5) // EF

	theGraph.Dfs()	
}
