package bstree

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
	return nil, true
}

func avlFind() {

}

func rightRotate(node *AvlNode) *AvlNode {
	return nil
}

func leftRotate(node *AvlNode) *AvlNode {
	rchild := node.right
	node.right = nil
	rchild.left = node
	return rchild
}

func height(node *AvlNode) int {
	if node == nil {
		return -1
	}
	return node.h
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
	node.h++
	lh := height(node.left)
	rh := height(node.right)
	diff := lh - rh
	if diff > 1 {
		if node.left != nil && node.left.left != nil {
			node = rightRotate(node.left)
		} else /*if node.left != nil && node.left.right != nil*/ {
			leftRotate(node.left)
			node = rightRotate(node)
		}
	} else if diff < -1 {
		//j
		if node.right != nil && node.right.right != nil {
			node = leftRotate(node)
		} else /*if node.right != nil && node.right.left != nil */ {
			rightRotate(node.right)
			node = leftRotate(node)
		}
	}
	return node
}

func (avl *AVLTree) Remove(key int) {

}
