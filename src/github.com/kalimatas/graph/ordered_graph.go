package graph

import (
	"fmt"
)

type OrderedGraph struct {
	Graph
}

func (g *OrderedGraph) AddEdge(start, end uint) (err error) {
	if start >= MAX_VERTS || end >= MAX_VERTS {
		err = &OutOfBoundariesError{"Vertex is out of bound"}
	}

	g.adjMatrix[start][end] = true

	return
}

// Topology sorting.
func (g *OrderedGraph) Topo() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("got an error: %s", err)
		}
	} ()

	for g.vertexCount > 0 {
		currentVertex := g.noSuccessors()
		if currentVertex == -1 {
			panic("Graph has cycles")
		}
		g.stack.Push(g.vertexList[currentVertex].Label)
		g.deleteVertex(uint(currentVertex))
	}

	fmt.Printf("Topology sorted: ")
	g.stack.Display()
	fmt.Println()
}

func (g *OrderedGraph) noSuccessors() int {
	var hasEdge bool

	for row := uint(0); row < g.vertexCount; row++ {
		hasEdge = false
		for col := uint(0); col < g.vertexCount; col++ {
			if g.adjMatrix[row][col] {
				hasEdge = true
				break
			}
		}

		if !hasEdge {
			return int(row);
		}
	}

	return -1
}

func (g *OrderedGraph) deleteVertex(delVertex uint) {
	if delVertex == g.vertexCount {
		return
	}

	for nVert := delVertex; nVert < g.vertexCount - 1; nVert++ {
		g.vertexList[nVert] = g.vertexList[nVert + 1]
	}

	for row := delVertex; row < g.vertexCount - 1; row++ {
		g.moveRowUp(row, g.vertexCount)
	}

	for col := delVertex; col < g.vertexCount - 1; col++ {
		g.moveColLeft(col, g.vertexCount - 1)
	}

	g.vertexCount--
}

func (g *OrderedGraph) moveRowUp(row, length uint) {
	for col := uint(0); col < length; col++ {
		g.adjMatrix[row][col] = g.adjMatrix[row + 1][col]
	}
}

func (g* OrderedGraph) moveColLeft(col, length uint) {
	for row := uint(0); row < length; row++ {
		g.adjMatrix[row][col] = g.adjMatrix[row][col + 1]
	}
}
