package graph

import "fmt"

type vertexT struct {
	vertex int
	weight int
}
type WeightGraph struct {
	vertexes []int
	edges    [][]vertexT
}

func (v *vertexT) Vertex() int {
	return v.vertex
}
func (v *vertexT) Weight() int {
	return v.weight
}

func NewWeightGraph(size int) *WeightGraph {
	return &WeightGraph{
		vertexes: make([]int, size),
		edges:    make([][]vertexT, size),
	}
}

func (g *WeightGraph) Size() int {
	return len(g.vertexes)
}

func (g *WeightGraph) AddEdge(start, end int, weight int) {
	g.vertexes[start] = start
	g.vertexes[end] = end

	g.edges[start] = append(g.edges[start], vertexT{vertex: end, weight: weight})
	g.edges[end] = append(g.edges[end], vertexT{vertex: start, weight: weight})
}

//顶点的临接顶点表
func (g *WeightGraph) Adjs(v int) []vertexT {
	return g.edges[v]
}

func (g *WeightGraph) Vertexes() []int {
	return g.vertexes
}
func (g *WeightGraph) Dump() {
	for i, edge := range g.edges {
		fmt.Printf("v:%d === ", i)
		for _, v := range edge {
			fmt.Printf(" %d", v.Vertex())
		}
		fmt.Printf("\n")
	}
}
