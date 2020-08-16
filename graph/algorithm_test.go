package graph

import (
	"testing"
)

//含有环的无向图 4-5-6 即环路径
var input = [][]int{
	{0, 4},
	{1, 3},
	{1, 4},
	{2, 5},
	{4, 5},
	{6, 5},
	{6, 4},
}

//含有两个连通分量的无向图
var inputConnectedPath = [][]int{
	{0, 4},
	{1, 3},
	{1, 4},
	{2, 5},
	//{4, 5},
	{6, 5},
	//{6, 4},
}

var dijkstraInput = [][]int{
	{1, 2, 2},
	{1, 3, 4},
	{1, 4, 1},

	{2, 4, 3},
	{2, 5, 10}, //5

	{3, 4, 2},
	{3, 6, 5},

	{4, 5, 7},
	{4, 6, 8},
	{4, 7, 4}, //10

	{5, 7, 6},

	{6, 7, 1},
}

func TestDFS(t *testing.T) {

}

func TestBFS(t *testing.T) {

}

/*

func TestIterator(t *testing.T) {
	size := len(input)
	g := NewGraph(size)
	for _, row := range input {
		g.AddEdge(row[0], row[1])
	}
	iter := Iterator(g, 1)
	for {
		if i, ok := iter(); ok {
			t.Logf("%d \n", i)
			continue
		}
		break
	}
}

func TestShortestSinglePath(t *testing.T) {
	size := len(input)
	g := NewGraph(size)
	for _, row := range input {
		g.AddEdge(row[0], row[1])
	}

	paths := SingleShortestPath(g, 0)

	start := 0
	for i := range paths {
		fmt.Printf("path: 0 -> %d   ", i)
		stack := make([]int, 0)
		next := i
		for {
			k := paths[next]
			stack = append(stack, k)
			if k == start {
				break
			}
			next = k
		}
		for j := len(stack) - 1; j >= 0; j-- {
			fmt.Printf(" %d", stack[j])
		}
		fmt.Printf("\n")
	}
}

func TestConnectedPath(t *testing.T) {
	size := len(input)
	g := NewGraph(size)
	for _, row := range inputConnectedPath {
		g.AddEdge(row[0], row[1])
	}

	count := getConnectedPathCount(g)
	if count != 2 {
		t.Fatalf("want 2 got:%d\n", count)
	}

	size = len(input)
	g = NewGraph(size)
	for _, row := range input {
		g.AddEdge(row[0], row[1])
	}
	count = getConnectedPathCount(g)
	if count != 1 {
		t.Fatalf("want 1 got:%d\n", count)
	}
}

func getConnectedPathCount(g Grapher) int {
	path := ConnectdPaths(g)
	count := 0

	for i, v := range path {
		if i == path[v] {
			count++
		}
	}
	return count
}

*/
func TestDijkstra(t *testing.T) {
	g := NewWeightGraph(8)
	for _, row := range dijkstraInput {
		g.AddEdge(row[0], row[1], row[2])
	}
	distances := dijkstra(g, 1)
	for i, v := range distances[1:] {
		t.Logf("vertex:%d distance:%d\n", i+1, v)
	}
}

func TestPrimMST(t *testing.T) {
	g := NewWeightGraph(8)
	for _, row := range dijkstraInput {
		g.AddEdge(row[0], row[1], row[2])
	}
	g.Dump()
	_, sum := PrimMST(g, 1)
	/*
		for i, v := range mst[1:] {
			t.Logf("vertex:%d end:%d\n", i+1, v)
		}
	*/
	t.Logf("sum:%d\n", sum)
}
