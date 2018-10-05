package misc

import (
	"container/list"
)

/*************************
 can use array to implement quik-find.
 however, find all the elements belong the same set will be harder
 when use this method
 0--1--2--3--4--5
 0--1--2--3--4--5

 union(2, 3)
 0--1--2--3--4--5
 0--1--2--2--4--5

 union(0, 5)
 0--1--2--3--4--5
 0--1--2--2--4--0

****************************/
type set struct {
	element  int
	elements *list.List
}

type QuickFind struct {
	sets []*set
}

func MakeSet(num int) *QuickFind {
	uset := &QuickFind{
		sets: make([]*set, num),
	}
	for i := 0; i < num; i++ {
		s := &set{element: i, elements: list.New()}
		s.elements.PushBack(i)
		uset.sets[i] = s
	}
	return uset
}

func (uset *QuickFind) IsSameSet(x, y int) bool {
	sets := uset.sets
	if sets[x].element == sets[y].element {
		return true
	}
	return false
}

func (uset *QuickFind) Find(element int) []int {
	elements := make([]int, 0)
	s := uset.sets[element]
	for e := s.elements.Front(); e != nil; e = e.Next() {
		elements = append(elements, e.Value.(int))
	}
	return elements
}

func (uset *QuickFind) Unite(src, dst int) {
	if src == dst {
		return
	}
	if src > dst {
		src, dst = dst, src
	}
	dstSet := uset.sets[dst]
	srcSet := uset.sets[src]

	for e := dstSet.elements.Front(); e != nil; e = e.Next() {
		i := e.Value.(int)
		srcSet.elements.PushBack(i)
		uset.sets[i] = srcSet
	}
	uset.sets[dst] = srcSet
}

func (uset *QuickFind) Dump() {
	/*
		for e, s := range uset.sets {
			fmt.Printf("element:%d  set:%v\n", e, s.elements)
		}
	*/
}

/*
type UnionFindSet struct {
	ids  []int
	size []int
}

func NewUnionFindSet(num int) *UnionFindSet {
	ufs := &UnionFindSet{
		ids:  make([]int, num, num),
		size: make([]int, 0, num),
	}
	for i := 0; i < num; i++ {
		ufs.ids[i] = i
	}
	return ufs
}

//TODO: path compression
func (ufs *UnionFindSet) root(i int) int {
	for ufs.ids[i] != i {
		i = ufs.ids[i]
	}
	return i
}

func (ufs *UnionFindSet) Find(x, y int) bool {
	return ufs.root(x) == ufs.root(y)
}

//TODO: path compression
func (ufs *UnionFindSet) Unite(x, y int) {
	rx := ufs.root(x)
	ry := ufs.root(y)

	ufs.ids[rx] = ry
}
*/
