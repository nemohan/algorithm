package graph

import (
	"testing"
)

func TestBFS(t *testing.T) {
	g := NewGraph()
	g.AddEdge(1, 2, 0)
	g.AddEdge(2, 3, 0)
	g.AddEdge(1, 4, 0)
	g.AddEdge(4, 3, 0)
	g.BFSTraverse()
	g.DFSTraverse()
	//g.Dump()
}

func TestShortestPath(t *testing.T) {
	g := NewGraph()
	g.AddEdge(65, 66, 3)
	g.AddEdge(65, 68, 7)
	g.AddEdge(66, 67, 4)
	g.AddEdge(66, 68, 2)
	g.AddEdge(67, 68, 5)
	g.AddEdge(67, 69, 6)
	g.AddEdge(68, 69, 4)
	g.ShortestPath(65)
}
func TestDFS(t *testing.T) {

}
