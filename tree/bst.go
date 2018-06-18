package tree

import (
	"fmt"
)

/******************
binary search tree

****************/

type bstNode struct {
	left   *bstNode
	right  *bstNode
	key    int
	value  interface{}
	height int
}

type BSTree struct {
	root *bstNode
	size int
}

func NewBSTree() *BSTree {
	return &BSTree{}
}

func (bst *BSTree) Insert(key int, value interface{}) {
	n, ok := find(&bst.root, key)
	if ok {
		return
	}
	*n = &bstNode{key: key, value: value}
	bst.size++
}

func (bst *BSTree) Find(key int) {
	n, ok := find(&bst.root, key)
	if !ok {
		return
	}
	fmt.Printf("key:%d\n", (*n).key)
}

// the pointer of pointer is better
//func find(root *bstNode, key int) (**bstNode, bool){
func find(root **bstNode, key int) (**bstNode, bool) {
	proot := *root
	if proot == nil {
		return root, false
	}
	if key == proot.key {
		return root, true
	}
	if key < proot.key {
		return find(&proot.left, key)
	}
	return find(&proot.right, key)
}

func (bst *BSTree) InOrder() {
	inOrder(bst.root)
}
func inOrder(root *bstNode) {
	if root == nil {
		return
	}
	inOrder(root.left)
	fmt.Printf("node:%d\n", root.key)
	inOrder(root.right)
}

func (bst *BSTree) PostOrder() {

}

func (bst *BSTree) PreOrder() {

}

func (bst *BSTree) BFSTraverse() {
	if bst.root == nil {
		return
	}
	queue := []*bstNode{bst.root}
	traverse(queue)
}

func traverse(queue []*bstNode) {
	if len(queue) == 0 {
		return
	}
	node := queue[0]
	queue = append(queue[:0], queue[1:]...)
	fmt.Printf("node:%v height:%d\n", node.key, node.height)
	if node.left != nil {
		queue = append(queue, node.left)
		node.left.height = node.height + 1
	}
	if node.right != nil {
		queue = append(queue, node.right)
		node.right.height = node.height + 1
	}
	traverse(queue)
}
