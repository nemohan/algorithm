package bstree

import (
	"fmt"
)

type AvlNode struct {
	left  *AvlNode
	right *AvlNode
	h     int
	key   int
	value interface{}
}

type AVLTree struct {
	root *AvlNode
}

func NewAVLTree() *AVLTree {
	return &AVLTree{
		root: nil,
	}
}

func (avl *AVLTree) Insert(key int, value interface{}) {
	avl.root = avlInsert(avl.root, key, value)
}

func (avl *AVLTree) Find(key int) (interface{}, bool) {
	node := avlFind(avl.root, key)
	if node == nil {
		return nil, false
	}
	return node.value, true
}

func avlFind(node *AvlNode, key int) *AvlNode {
	if node == nil {
		return nil
	}
	if key < node.key {
		return avlFind(node.left, key)
	} else if key > node.key {
		return avlFind(node.right, key)
	}
	return node
}

func rightRotate(node *AvlNode) *AvlNode {
	lchild := node.left
	node.left = lchild.right
	lchild.right = node

	node.h = max(height(node.left), height(node.right)) + 1
	lchild.h = max(height(lchild.left), height(lchild.right)) + 1
	return lchild
}

func leftRotate(node *AvlNode) *AvlNode {
	rchild := node.right
	node.right = rchild.left
	rchild.left = node

	node.h = max(height(node.left), height(node.right)) + 1
	rchild.h = max(height(rchild.left), height(rchild.right)) + 1
	return rchild
}

func height(node *AvlNode) int {
	if node == nil {
		return -1
	}
	return node.h
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func avlInsert(node *AvlNode, key int, value interface{}) *AvlNode {
	if node == nil {
		return &AvlNode{h: 0, key: key, value: value}
	}
	if key < node.key {
		node.left = avlInsert(node.left, key, value)
	} else if key > node.key {
		node.right = avlInsert(node.right, key, value)
	}
	lh := height(node.left)
	rh := height(node.right)
	diff := lh - rh
	if diff > 1 {
		if key < node.left.key {
			node = rightRotate(node)
		} else if key > node.left.key {
			node.left = leftRotate(node.left)
			node = rightRotate(node)
		}
	} else if diff < -1 {
		if key > node.right.key {
			node = leftRotate(node)
		} else if key < node.right.key {
			node.right = rightRotate(node.right)
			node = leftRotate(node)
		}
	}
	node.h = max(height(node.left), height(node.right)) + 1
	return node
}

func (avl *AVLTree) Delete(key int) {
	avl.root = avlDelete(avl.root, key)
}

func findMin(node *AvlNode) *AvlNode {
	if node.left == nil {
		return node
	}
	return findMin(node.left)
}

func avlDelete(node *AvlNode, key int) *AvlNode {
	if node == nil {
		return nil
	}
	if key < node.key {
		node.left = avlDelete(node.left, key)
	} else if key > node.key {
		node.right = avlDelete(node.right, key)
	} else {
		if node.left == nil && node.right == nil {
			return nil
		} else if node.left != nil && node.right == nil {
			return node.left
		} else if node.right != nil && node.left == nil {
			return node.right
		} else { //two subtree
			min := findMin(node.right)
			node.key = min.key
			node.value = min.value
			node.right = avlDelete(node.right, min.key)
		}
	}

	lh := height(node.left)
	rh := height(node.right)
	diff := lh - rh
	if diff > 1 {
		if node.left != nil && node.left.left != nil {
			node = rightRotate(node)
		} else if node.left != nil && node.left.right != nil {
			node.left = leftRotate(node.left)
			node = rightRotate(node)
		}
	} else if diff < -1 {
		if node.right != nil && node.right.right != nil {
			node = leftRotate(node)
		} else if node.right != nil && node.right.left != nil {
			node.right = rightRotate(node.right)
			node = leftRotate(node)
		}
	}
	node.h = max(height(node.left), height(node.right)) + 1
	return node
}

func (avl *AVLTree) Dump() {
	dumpTree(avl.root)
}
func (avl *AvlNode) Left() Node {
	return avl.left
}

func (avl *AvlNode) Right() Node {
	return avl.right
}

func (avl *AvlNode) Key() int {
	return avl.key
}

func (avl *AvlNode) String() string {
	return fmt.Sprintf("{key:%d height:%d}\n", avl.key, avl.h)
}
