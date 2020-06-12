package bstree

import (
	"fmt"
	"testing"
)

/*
   test with go test -run="Splay"
*/

func insertSTNode(node *splayNode, k int, v interface{}) *splayNode {
	if node == nil {
		return &splayNode{key: k, value: v}
	}
	if k < node.key {
		node.left = insertSTNode(node.left, k, v)
	} else if k > node.key {
		node.right = insertSTNode(node.right, k, v)
	}
	return node
}

func TestSplayTreeFind(t *testing.T) {
	stTree := NewSplayTree()
	for i := 7; i >= 1; i-- {
		stTree.root = insertSTNode(stTree.root, i, nil)
	}
	dumpTree(stTree.root)

	if _, ok := stTree.Find(1); !ok {
		dumpTree(stTree.root)
		t.Fatal("not found valued 1\n")
	}
	fmt.Printf("after splay\n")
	dumpTree(stTree.root)
}

func TestSplayTree(t *testing.T) {
	stTree := NewSplayTree()
	for i := 4; i >= 1; i-- {
		stTree.root = insertSTNode(stTree.root, i, nil)
	}
	if _, ok := stTree.Find(1); !ok {
		dumpTree(stTree.root)
		t.Error("not found valued 1\n")
	}
	dumpTree(stTree.root)
}
