package graph

import (
	"testing"
)

func TestBFS(t *testing.T) {
	g := NewGraph()
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(1, 4)
	g.AddEdge(4, 3)
	g.BFSTraverse()
	g.DFSTraverse()
	//g.Dump()
}

func TestDFS(t *testing.T) {

}
