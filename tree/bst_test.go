package tree

import (
	"testing"
)

func createTree() *BSTree {
	bst := NewBSTree()
	bst.Insert(2, 2)
	bst.Insert(3, 3)
	bst.Insert(1, 1)
	return bst
}
func TestInsert(t *testing.T) {
	bst := createTree()
	bst.Find(1)
	bst.BFSTraverse()
}

func TestInOrder(t *testing.T) {
	bst := createTree()
	bst.Insert(8, 0)
	bst.Insert(10, 0)
	bst.Insert(6, 0)
	bst.InOrder()
}
