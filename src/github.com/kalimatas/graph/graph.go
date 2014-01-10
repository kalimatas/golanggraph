package graph

import (
	"github.com/kalimatas/structures"
)

const MAX_VERTS uint = 20

type Graph struct {
	VertexList  structures.List
	AdjMatrix   [MAX_VERTS][MAX_VERTS]bool // adjacency matrix
	VertexCount uint                       // current number of verteces
}

type OutOfBoundariesError struct {
	message string
}

func (e *OutOfBoundariesError) Error() string {
	return e.message
}

func (graph *Graph) addVertex(vertex *Vertex) (err error) {
	if graph.VertexCount >= MAX_VERTS {
		err = &OutOfBoundariesError{"Reached max verteces count"}
	}

	graph.VertexList.InsertFirst(vertex)
	graph.VertexCount++

	return
}

func (graph *Graph) addEdge(start, end uint) (err error) {
	if start >= MAX_VERTS || end >= MAX_VERTS {
		err = &OutOfBoundariesError{"Vertex is out of bound"}
	}

	graph.AdjMatrix[start][end] = true
	graph.AdjMatrix[end][start] = true

	return
}
