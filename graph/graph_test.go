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
func TestShortestPath2(t *testing.T) {
	g := NewGraph()
	g.AddEdge(65, 66, 3)
	g.AddEdge(65, 67, 5)
	g.AddEdge(65, 68, 4)

	g.AddEdge(66, 69, 3)
	g.AddEdge(66, 70, 6) //5

	g.AddEdge(67, 68, 2)
	g.AddEdge(67, 71, 4)

	g.AddEdge(68, 69, 1)
	g.AddEdge(68, 72, 5)

	g.AddEdge(69, 70, 2) // 10
	g.AddEdge(69, 73, 4)

	g.AddEdge(70, 74, 5)

	g.AddEdge(71, 72, 3)
	g.AddEdge(71, 75, 6)

	g.AddEdge(72, 73, 6) //5
	g.AddEdge(72, 75, 7)

	g.AddEdge(73, 74, 3)
	g.AddEdge(73, 76, 5)

	g.AddEdge(74, 76, 9)
	g.AddEdge(75, 76, 8)
	g.ShortestPath(65)
}

func TestDFS(t *testing.T) {

}
