package graph

import (
	"container/list"
	"fmt"
)

type vertex struct {
	id     int
	weight int
	list   *list.List
}

type Graph struct {
	vertexes map[int]*vertex
	visited  map[int]bool
}

func NewGraph() *Graph {
	return &Graph{
		vertexes: make(map[int]*vertex),
		visited:  make(map[int]bool),
	}
}

func (g *Graph) addEdge(head int, tail int) {
	v, ok := g.vertexes[head]
	if !ok {
		v = &vertex{id: head, list: list.New()}
		g.vertexes[head] = v
	}
	v.list.PushBack(tail)
}

//no direction
func (g *Graph) AddEdge(head int, tail int) {
	g.addEdge(head, tail)
	g.addEdge(tail, head)
}

func (g *Graph) BFSTraverse() {
	queue := list.New()
	for _, v := range g.vertexes {
		if _, ok := g.visited[v.id]; ok {
			continue
		}
		queue.PushBack(v)
		g.bfs(queue)
	}
}

func (g *Graph) bfs(queue *list.List) {
	if queue.Len() == 0 {
		return
	}
	head := queue.Front()
	queue.Remove(head)
	v := head.Value.(*vertex)
	l := v.list
	fmt.Printf("bfs vertex:%d\n", v.id)
	g.visited[v.id] = true
	for e := l.Front(); e != nil; e = e.Next() {
		id := e.Value.(int)
		if _, ok := g.visited[id]; ok {
			continue
		}
		g.visited[id] = true
		queue.PushBack(g.vertexes[id])
	}
	g.bfs(queue)
}

func DFSTraverse() {

}
func dfs() {

}

func (g *Graph) Dump() {
	for i, v := range g.vertexes {
		fmt.Printf("head:%d\n", i)
		l := v.list
		for e := l.Front(); e != nil; e = e.Next() {
			id := e.Value.(int)
			fmt.Printf("%d, ", id)
		}
		fmt.Printf("\n-------------\n")
	}
}
