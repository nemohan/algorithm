package bstree

import (
	"fmt"
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

func (n *BSTNode) Left() Node {
	return n.left
}
func (n *BSTNode) Right() Node {
	return n.right
}
func (n *BSTNode) Key() int {
	return n.k
}
func (n *BSTNode) String() string {
	return fmt.Sprintf("{key:%d}\n", n.k)
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
	if key < node.k {
		return find(node.left, key)
	} else if key > node.k {
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
	} else {
		node.v = v
	}
	return node
}

func (b *BSTree) Delete(key int) {
	b.root = deleteBSTNode(b.root, key)
}

func deleteBSTNode(node *BSTNode, key int) *BSTNode {
	if node == nil {
		return nil
	}
	if key < node.k {
		node.left = deleteBSTNode(node.left, key)
	} else if key > node.k {
		node.right = deleteBSTNode(node.right, key)
	} else {
		return delNode(node)
	}
	return node
}

func delNode(node *BSTNode) *BSTNode {
	if node.left == nil && node.right == nil {
		return nil
	} else if node.left == nil && node.right != nil {
		return node.right
	} else if node.right == nil && node.left != nil {
		return node.left
	}
	//node with two subtrees
	minNode := findMinBstNode(node.right)
	node.k = minNode.k
	node.v = minNode.v
	node.right = deleteBSTNode(node.right, minNode.k)
	/*
		p := node.right
		parent := node
		for p != nil {
			if p.left == nil {
				break
			}
			parent = p
			p = p.left
		}
		node.k = p.k
		if p.right != nil {
			parent.left = p.right
		}
		if parent == node {
			parent.right = nil
		}
	*/
	return node
}

func findMinBstNode(node *BSTNode) *BSTNode {
	if node.left == nil {
		return node
	}
	return findMinBstNode(node.left)
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

/*
func (b *BSTree) Dump() {
	if b.root == nil {
		return
	}
	level := 0
	dump(b.root, level)
}
*/
