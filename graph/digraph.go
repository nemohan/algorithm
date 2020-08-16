package graph

type DiGraph struct {
	vertexes []int
	edges    [][]int
}

func NewDiGraph(size int) *DiGraph {
	return &DiGraph{
		vertexes: make([]int, size),
		edges:    make([][]int, size),
	}
}

func (g *DiGraph) Size() int {
	return len(g.vertexes)
}

func (g *DiGraph) AddEdge(start, end int) {
	g.vertexes[start] = start
	g.vertexes[end] = end

	g.edges[start] = append(g.edges[start], end)
	g.edges[end] = append(g.edges[end], start)
}

//顶点的临接顶点表
func (g *DiGraph) Adjs(v int) []int {
	return g.edges[v]
}

func (g *DiGraph) Vertexes() []int {
	return g.vertexes
}
