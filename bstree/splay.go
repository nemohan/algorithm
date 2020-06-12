package bstree

import (
	"fmt"
)

type splayNode struct {
	key   int
	value interface{}
	left  *splayNode
	right *splayNode
}

type SplayTree struct {
	root *splayNode
}

func NewSplayTree() *SplayTree {
	return &SplayTree{}
}

func (s *SplayTree) Insert() {

}

func (s *SplayTree) Find(key int) (interface{}, bool) {
	if s.root == nil {
		return nil, false
	}
	s.root = splayTreeFind(s.root, key)
	if s.root.key == key {
		return s.root.value, true
	} else if s.root.left != nil && s.root.left.key == key {
		s.root = rightRotateST(s.root)
	} else if s.root.right != nil && s.root.right.key == key {
		s.root = leftRotateST(s.root)
	} else {
		return nil, false
	}
	return s.root.value, true
}

func rightRotateST(node *splayNode) *splayNode {
	l := node.left
	rchild := l.right
	l.right = node
	node.left = rchild
	return l
}

func leftRotateST(node *splayNode) *splayNode {
	r := node.right
	lchild := r.left
	r.left = node
	node.right = lchild
	return r
}

func zigZigLeft(node *splayNode) *splayNode {
	fmt.Printf("zigzigLeft\n")
	r := node.right
	leftRotateST(node)
	node = leftRotateST(r)
	return node
}

func zigZigRight(node *splayNode) *splayNode {
	fmt.Printf("zigzigRight\n")
	l := node.left
	rightRotateST(node)
	node = rightRotateST(l)
	return node
}

func zigZagRL(node *splayNode) *splayNode {
	fmt.Printf("zigZagRL\n")
	node.right = rightRotateST(node.right)
	return leftRotateST(node)
}

//zigZagLeft left right
func zigZagLR(node *splayNode) *splayNode {
	fmt.Printf("zigZagLR\n")
	node.left = leftRotateST(node.left)
	return rightRotateST(node)
}

func splayTreeFind(node *splayNode, key int) *splayNode {
	if node == nil {
		return nil
	}
	if key < node.key {
		node.left = splayTreeFind(node.left, key)
		if node.left != nil && node.left.left != nil && node.left.left.key == key {
			node = zigZigRight(node)
		} else if node.left != nil && node.left.right != nil && node.left.right.key == key {
			node = zigZagRL(node)
		}
	} else if key > node.key {
		node.right = splayTreeFind(node.right, key)
		if node.right != nil && node.right.right != nil && node.right.right.key == key {
			node = zigZigLeft(node)
		} else if node.right != nil && node.right.left != nil && node.right.right.key == key {
			node = zigZagLR(node)
		}
	}
	return node
}

func (s *SplayTree) Delete() {

}

func (n *splayNode) Left() Node {
	return n.left
}
func (n *splayNode) Right() Node {
	return n.right
}
func (n *splayNode) Key() int {
	return n.key
}
func (n *splayNode) String() string {
	return fmt.Sprintf("{key:%d}\n", n.key)
}
