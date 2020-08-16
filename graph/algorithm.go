package graph

import (
	"fmt"
	"math"
)

/*
//从某个顶点开始遍历
func dfs(g Grapher, start int, marked []bool) {
	marked[start] = true
	for _, v := range g.Adjs(start) {
		if !marked[v] {
			dfs(g, v, marked)
		}
	}
}

//广度优先遍历
func bfs(g Grapher, start int) {
	queue := []int{start}
	marked := make([]bool, g.Size())
	marked[start] = true
	for len(queue) > 0 {
		v := queue[0]
		for _, k := range g.Adjs(v) {
			if !marked[k] {
				marked[k] = true
				queue = append(queue, k)
			}
		}
		queue = queue[1:]
	}
}



//连通分量
func ConnectdPaths(g Grapher) []int {
	edgeTo := make([]int, g.Size())
	marked := make([]bool, g.Size())
	for _, v := range g.Vertexes() {
		if !marked[v] {
			edgeTo[v] = v
		}
		dfsConnected(g, v, edgeTo, marked)
	}
	return edgeTo
}

func dfsConnected(g Grapher, start int, edgeTo []int, marked []bool) {
	marked[start] = true
	for _, v := range g.Adjs(start) {
		if !marked[v] {
			dfsConnected(g, v, edgeTo, marked)
			edgeTo[v] = start
		}
	}
}

//无权单源最短路径
func SingleShortestPath(g Grapher, start int) []int {
	queue := []int{start}
	marked := make([]bool, g.Size())
	edgeTo := make([]int, g.Size())
	marked[start] = true
	for len(queue) > 0 {
		v := queue[0]
		for _, k := range g.Adjs(v) {
			if !marked[k] {
				edgeTo[k] = v
				marked[k] = true
				queue = append(queue, k)
			}
		}
		queue = queue[1:]
	}
	return edgeTo
}

*/
//最小生成树的prim算法
func PrimMST(g Grapher, start int) ([]int, int) {
	queue := []int{start}
	size := g.Size()
	distances := make([]int, size)
	marked := make([]bool, size)
	for i := range distances {
		distances[i] = int(math.MaxInt32)
	}

	distances[start] = 0
	for len(queue) > 0 {
		s := queue[0]
		marked[s] = true
		for _, vertex := range g.Adjs(s) {
			v := vertex.Vertex()
			if !marked[v] && vertex.Weight() < distances[v] {
				distances[v] = vertex.Weight()
			}
		}
		min := int(math.MaxInt32)
		minIndex := 0
		for j, d := range distances {
			if !marked[j] && d < min {
				min = d
				minIndex = j
			}
		}
		fmt.Printf("min vertex:%d\n", minIndex)
		if !marked[minIndex] {
			queue = append(queue, minIndex)
		}
		queue = queue[1:]
	}
	sum := 0
	for i, d := range distances[1:] {
		fmt.Printf("%d distance:%d\n", i+1, d)
		sum += d
	}
	return nil, sum
}

func dijkstra(g Grapher, start int) []int {
	marked := make([]bool, g.Size())
	queue := make([]int, 0)
	distances := make([]int, g.Size())
	for i := range distances {
		distances[i] = int(math.MaxInt32)
	}
	distances[start] = 0
	queue = append(queue, start)
	for len(queue) > 0 {
		s := queue[0]
		marked[s] = true
		for _, vertex := range g.Adjs(s) {
			v := vertex.Vertex()
			d := distances[s] + vertex.Weight()
			if d < distances[v] {
				distances[v] = d
			}
			if !marked[v] {
				queue = append(queue, v)
			}
		}
		queue = queue[1:]
	}
	return distances
}
