package graph

const MAX_VERTS uint = 20

type Graph struct {
	vertexList  [MAX_VERTS]*Vertex
	adjMatrix   [MAX_VERTS][MAX_VERTS]bool // adjacency matrix
	VertexCount uint                       // current number of verteces
}

type OutOfBoundariesError struct {
	message string
}

func (e *OutOfBoundariesError) Error() string {
	return e.message
}

func (graph *Graph) AddVertex(vertex *Vertex) (err error) {
	if graph.VertexCount >= MAX_VERTS {
		err = &OutOfBoundariesError{"Reached max verteces count"}
	}

	graph.VertexCount++
	graph.vertexList[graph.VertexCount] = vertex
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
	return graph.VertexCount
}