package bstree

import (
	"testing"
)

var data = []int{1, 4, 5, 2, 3}

func TestFind(t *testing.T) {

}

func TestInsert(t *testing.T) {
	tree := NewBSTree()
	for _, v := range data {
		tree.Insert(v, nil)
	}
	tree.InOrderTraverse()
}

func TestDump(t *testing.T) {
	tree := NewBSTree()
	for _, v := range data {
		tree.Insert(v, nil)
	}
	tree.Dump()
}
func BenchmarkFind(t *testing.B) {

}
