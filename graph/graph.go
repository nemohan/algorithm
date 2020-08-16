package graph

type Grapher interface {
	AddEdge(start int, end int, weight int)
	Adjs(v int) []vertexT
	Size() int
	Vertexes() []int
}

/*
func NewGraph(size int) *Graph {
	return &Graph{
		vertexes: make([]int, size),
		edges:    make([][]int, size),
	}
}

type Graph struct {
	vertexes []int
	edges    [][]int
}

func (g *Graph) Size() int {
	return len(g.vertexes)
}

func (g *Graph) AddEdge(start, end int) {
	g.vertexes[start] = start
	g.vertexes[end] = end

	g.edges[start] = append(g.edges[start], end)
	g.edges[end] = append(g.edges[end], start)
}

//顶点的临接顶点表
func (g *Graph) Adjs(v int) []int {
	return g.edges[v]
}

func (g *Graph) Vertexes() []int {
	return g.vertexes
}

func (g *Graph) Dump() {
	for i, edge := range g.edges {
		fmt.Printf("v:%d === ", i)
		for _, v := range edge {
			fmt.Printf(" %d", v)
		}
		fmt.Printf("\n")
	}
}

func Iterator(g Grapher, start int) func() (int, bool) {
	i := 0
	return func() (int, bool) {
		vertexes := g.Adjs(start)
		v := -1
		if i < len(vertexes) {
			v := vertexes[i]
			i++
			return v, true
		}
		return v, false
	}
}

*/
