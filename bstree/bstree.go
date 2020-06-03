package bstree

import (
	"fmt"
	"reflect"
)

type BSTNode struct {
	v     interface{}
	k     int
	left  *BSTNode
	right *BSTNode
}

type BSTree struct {
	root *BSTNode
}

func NewBSTree() *BSTree {
	return &BSTree{
		root: nil,
	}
}

func (b *BSTree) Find(key int) (interface{}, bool) {
	node := find(b.root, key)
	if node == nil {
		return nil, false
	}
	return node.v, true
}

func find(node *BSTNode, key int) *BSTNode {
	if node == nil {
		return nil
	}
	if node.k < key {
		return find(node.left, key)
	} else if node.k > key {
		return find(node.right, key)
	}
	return node
}

func (b *BSTree) Insert(k int, v interface{}) {
	b.root = insert(b.root, k, v)
}

func insert(node *BSTNode, k int, v interface{}) *BSTNode {
	if node == nil {
		return &BSTNode{k: k, v: v}
	}
	if k < node.k {
		node.left = insert(node.left, k, v)
	} else if k > node.k {
		node.right = insert(node.right, k, v)
	}
	return node
}

func (b *BSTree) Delete() {

}

func (b *BSTree) InOrderTraverse() {
	inOrderTraverse(b.root)
}

func inOrderTraverse(node *BSTNode) {
	if node == nil {
		return
	}
	inOrderTraverse(node.left)
	fmt.Printf("%d\n", node.k)
	inOrderTraverse(node.right)
}

func (b *BSTree) Dump() {
	if b.root == nil {
		return
	}
	level := 0
	dump(b.root, level)
}

func (n *BSTNode) Left() Node {
	return n.left
}
func (n *BSTNode) Right() Node {
	return n.right
}
func (n *BSTNode) Key() int {
	return n.k
}

func dump(node Node, level int) {
	if node == nil {
		return
	}
	queue := []Node{node}
	fmt.Printf("level:%d %d\n", level, node.Key())
	for len(queue) > 0 {
		level++
		node := queue[0]
		left := node.Left()
		right := node.Right()
		lv := reflect.ValueOf(left)
		rv := reflect.ValueOf(right)
		if left != nil && !lv.IsNil() {
			fmt.Printf("left level:%d %d\n", level, left.Key())
			queue = append(queue, left)
		}
		if right != nil && !rv.IsNil() {
			fmt.Printf("right level:%d %d\n", level, right.Key())
			queue = append(queue, right)
		}
		queue = queue[1:]
	}
}
