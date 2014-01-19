package main

import (
	"fmt"
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

	fmt.Printf("Unordered graph\n")
	fmt.Println("Depth-first search")
	theGraph.Dfs()	
	fmt.Println()
	
	fmt.Println("Breadth-first search")
	theGraph.Bfs()
	fmt.Println()

	fmt.Printf("Ordered graph\n\n")
	orderedGraph := new(graph.OrderedGraph)
	orderedGraph.AddVertex(&graph.Vertex{Label: "A"}) // 0
	orderedGraph.AddVertex(&graph.Vertex{Label: "B"}) // 1
	orderedGraph.AddVertex(&graph.Vertex{Label: "C"}) // 2
	orderedGraph.AddVertex(&graph.Vertex{Label: "D"}) // 3
	orderedGraph.AddVertex(&graph.Vertex{Label: "E"}) // 4
	orderedGraph.AddVertex(&graph.Vertex{Label: "F"}) // 5
	orderedGraph.AddVertex(&graph.Vertex{Label: "G"}) // 6
	orderedGraph.AddVertex(&graph.Vertex{Label: "H"}) // 7

	orderedGraph.AddEdge(0, 3)
	orderedGraph.AddEdge(0, 4)
	orderedGraph.AddEdge(1, 4)
	orderedGraph.AddEdge(2, 5)
	orderedGraph.AddEdge(3, 6)
	orderedGraph.AddEdge(4, 6)
	orderedGraph.AddEdge(5, 7)
	orderedGraph.AddEdge(6, 7)

	orderedGraph.Topo()
}
