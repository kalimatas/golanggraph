package graph

import (
	"fmt"
	"github.com/kalimatas/structures"
)

const MAX_VERTS uint = 20

type Graph struct {
	vertexList  [MAX_VERTS]*Vertex
	adjMatrix   [MAX_VERTS][MAX_VERTS]bool // adjacency matrix
	vertexCount uint                       // current number of verteces
	stack       structures.Stack           // for Dfs
	queue       structures.Queue           // for Bfs
}

type OutOfBoundariesError struct {
	message string
}

func (e *OutOfBoundariesError) Error() string {
	return e.message
}

func (graph *Graph) AddVertex(vertex *Vertex) (err error) {
	if graph.vertexCount >= MAX_VERTS {
		// todo: expand array
		err = &OutOfBoundariesError{"Reached max verteces count"}
	}

	graph.vertexList[graph.vertexCount] = vertex
	graph.vertexCount++
	return
}

func (graph *Graph) AddEdge(start, end uint) (err error) {
	if start >= MAX_VERTS || end >= MAX_VERTS {
		err = &OutOfBoundariesError{"Vertex is out of bound"}
	}

	graph.adjMatrix[start][end] = true
	graph.adjMatrix[end][start] = true

	return
}

func (graph *Graph) GetVertextCount() (count uint) {
	return graph.vertexCount
}

// Depth-first search
func (graph *Graph) Dfs() {
	if graph.vertexCount == 0 {
		return
	}

	graph.vertexList[0].IsVisited = true
	fmt.Printf("%s ", graph.vertexList[0])
	graph.stack.Push(uint(0))

	for !graph.stack.IsEmpty() {
		adjVertex := graph.getAdjUnvisitedVertex(graph.stack.Peek().(uint))
		if adjVertex == -1 {
			graph.stack.Pop()
		} else {
			graph.vertexList[adjVertex].IsVisited = true
			fmt.Printf("%s ", graph.vertexList[adjVertex])
			graph.stack.Push(uint(adjVertex))
		}
	}

	// reset flags
	for i := uint(0); i < graph.vertexCount; i++ {
		graph.vertexList[i].IsVisited = false
	}
}

func (graph *Graph) Bfs() {
	if graph.vertexCount == 0 {
		return
	}

	graph.vertexList[0].IsVisited = true
	fmt.Printf("%s ", graph.vertexList[0])
	graph.queue.Insert(uint(0))
	
	for !graph.queue.IsEmpty() {
		currentVertex := graph.queue.Remove().(uint)
		
		for nextVertex := graph.getAdjUnvisitedVertex(currentVertex); 
			nextVertex != -1;
			nextVertex = graph.getAdjUnvisitedVertex(currentVertex)	{
			
			graph.vertexList[nextVertex].IsVisited = true
			fmt.Printf("%s ", graph.vertexList[nextVertex])
			graph.queue.Insert(uint(nextVertex))
		}			
	}
	
	// reset flags
	for i := uint(0); i < graph.vertexCount; i++ {
		graph.vertexList[i].IsVisited = false
	}
}

// Returns not visited adjacent vertex.
func (graph *Graph) getAdjUnvisitedVertex(vertex uint) (unvisited int) {
	for i := uint(0); i < graph.vertexCount; i++ {
		if graph.adjMatrix[vertex][i] == true && graph.vertexList[i].IsVisited == false {
			return int(i)
		}
	}

	return -1
}
