package graph

import (
	"container/list"
	"fmt"
)

type vertex struct {
	id       int
	list     *list.List
	next     *vertex
	prev     *vertex
	distance int
	key      interface{}
	prevID   int
}

type edge struct {
	end    int
	start  int
	weight int
}

type Graph struct {
	vertexes map[int]*vertex
	visited  map[int]bool
}

const maxWeight = 100

func NewGraph() *Graph {
	return &Graph{
		vertexes: make(map[int]*vertex),
		visited:  make(map[int]bool),
	}
}

func (g *Graph) AddVertex(id int) {
	_, ok := g.vertexes[id]
	if ok {
		return
	}
	g.vertexes[id] = &vertex{id: id}
}

func (g *Graph) addEdge(head int, tail int, weight int) {
	v, ok := g.vertexes[head]
	if !ok {
		v = &vertex{id: head, list: list.New()}
		g.vertexes[head] = v
	}
	v.list.PushBack(&edge{end: tail, weight: weight})
}

//no direction
func (g *Graph) AddEdge(head int, tail int, weight int) {
	g.addEdge(head, tail, weight)
	g.addEdge(tail, head, weight)
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
		pe := e.Value.(*edge)
		if _, ok := g.visited[pe.end]; ok {
			continue
		}
		g.visited[pe.end] = true
		queue.PushBack(g.vertexes[pe.end])
	}
	g.bfs(queue)
}

func (g *Graph) DFSTraverse() {
	g.visited = make(map[int]bool)
	for _, v := range g.vertexes {
		if _, ok := g.visited[v.id]; ok {
			continue
		}
		g.dfs(v)
	}
}

func (g *Graph) dfs(v *vertex) {
	fmt.Printf("dfs vertex:%d\n", v.id)
	g.visited[v.id] = true
	l := v.list
	for e := l.Front(); e != nil; e = e.Next() {
		pe := e.Value.(*edge)
		if _, ok := g.visited[pe.end]; ok {
			continue
		}

		g.dfs(g.vertexes[pe.end])
	}
}

func chooseMinDistance(vertexSlice []*vertex) ([]*vertex, *vertex) {
	min := 0
	for i := 1; i < len(vertexSlice); i++ {
		if vertexSlice[i].distance < vertexSlice[min].distance {
			min = i
		}
	}
	v := vertexSlice[min]
	vertexSlice = append(vertexSlice[:min], vertexSlice[min+1:]...)
	return vertexSlice, v
}

//ShortestPath use Dijsktra
func (g *Graph) ShortestPath(start int) {
	vertexSet := make([]*vertex, 0)
	for _, v := range g.vertexes {
		v.distance = maxWeight
		vertexSet = append(vertexSet, v)
	}
	pathTable := make(map[int]*vertex)
	g.vertexes[start].distance = 0
	for i := 0; i < len(g.vertexes); i++ {
		var v *vertex
		vertexSet, v = chooseMinDistance(vertexSet)
		pathTable[v.id] = v
		if v.prev != nil {
			v.prev.next = v
		}
		fmt.Printf("vertex %d\n", v.id)
		for e := v.list.Front(); e != nil; e = e.Next() {
			pe := e.Value.(*edge)
			id := pe.end
			t := g.vertexes[id]
			fmt.Printf("v:%d v2:%d dis:%d weight:%d\n", v.id, id, v.distance, pe.weight)
			if id != v.id && v.distance+pe.weight < t.distance {
				t.distance = v.distance + pe.weight
				t.prev = v
			}
		}
	}

	for _, v := range pathTable {
		fmt.Printf("%d distance:%d\n", v.id, v.distance)
	}
}

func PrimSpanningTree() {

}

func KruskalSpanningTree() {

}

func chooseMinEdge(edgeSet map[int]*edge) *edge {
	min := maxWeight
	i := 0
	for k, e := range edgeSet {
		if e.weight < min {
			i = k
			min = e.weight
		}
	}
	e := edgeSet[i]
	delete(edgeSet, i)
	return e
}

func (g *Graph) GenSpanningTree() {
	edgeSet := make(map[int]*edge, 0)
	lastVertex := (*vertex)(nil)
	for _, v := range g.vertexes {
		e := &edge{end: v.id, weight: maxWeight}
		edgeSet[v.id] = e
		lastVertex = v
	}

	//current := g.vertexes[0]
	current := lastVertex
	delete(edgeSet, current.id)
	tree := make(map[int]*edge)
	tree[current.id] = nil

	for len(edgeSet) > 0 {
		for e := current.list.Front(); e != nil; e = e.Next() {
			pe := e.Value.(*edge)
			_, ok := tree[pe.end]
			fmt.Printf("end:%d \n", pe.end)
			if !ok && pe.weight < edgeSet[pe.end].weight {
				edgeSet[pe.end].weight = pe.weight
				edgeSet[pe.end].start = current.id
			}
		}
		minEdge := chooseMinEdge(edgeSet)
		fmt.Printf("add start:%d end:%d weight:%d\n", minEdge.start, minEdge.end, minEdge.weight)
		tree[minEdge.end] = minEdge
		current = g.vertexes[minEdge.end]
	}
	fmt.Printf("tree:%v\n", tree)
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
