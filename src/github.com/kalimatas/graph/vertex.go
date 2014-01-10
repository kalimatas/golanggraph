package graph

import "fmt"

type Vertex struct {
	Label     string
	IsVisited bool
}

func (vertex *Vertex) String() string {
	return fmt.Sprintf("vertex '%s'", vertex.Label)
}
