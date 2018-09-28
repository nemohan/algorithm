package misc

import (
	"fmt"
)

type set struct {
	elements map[int]bool
}

type UnionFindSet struct {
	sets map[int]*set
}

func MakeSet(src []int) *UnionFindSet {
	uset := &UnionFindSet{
		sets: make(map[int]*set, 0),
	}
	for _, e := range src {
		s := &set{elements: make(map[int]bool, 0)}
		s.elements[e] = false
		uset.sets[e] = s
	}
	return uset
}

func (uset *UnionFindSet) Find(element int) []int {
	elements := make([]int, 0)
	s := uset.sets[element]
	for k, _ := range s.elements {
		elements = append(elements, k)
	}
	return elements
}

func (uset *UnionFindSet) Union(src, dst int) {
	srcSet := uset.sets[src]
	dstSet := uset.sets[dst]
	for k, _ := range dstSet.elements {
		srcSet.elements[k] = false
		uset.sets[k] = srcSet
	}
}

func (uset *UnionFindSet) Dump() {
	for e, s := range uset.sets {
		fmt.Printf("element:%d  set:%v\n", e, s.elements)
	}
}
