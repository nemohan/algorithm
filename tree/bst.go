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
	parent *bstNode
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
		(*n).value = value
		return
	}
	*n = &bstNode{key: key, value: value}
	bst.size++
}

func (bst *BSTree) Find(key int) (interface{}, bool) {
	n, ok := find(&bst.root, key)
	if !ok {
		return nil, false
	}
	return (*n).key, true
}

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

func (bst *BSTree) Del(key int) bool {
	if bst.root == nil {
		return true
	}
	node, ok := find(&(bst.root), key)
	if !ok {
		return true
	}
	dstNode := *node
	if dstNode.left == nil && dstNode.right == nil {
		*node = nil
		return true
	}
	if dstNode.left == nil && dstNode.right != nil {
		dstNode.key = dstNode.right.key
		dstNode.value = dstNode.right.value
		dstNode.right = nil
		return true
	} else if dstNode.left != nil && dstNode.right == nil {
		dstNode.key = dstNode.left.key
		dstNode.value = dstNode.left.value
		dstNode.left = nil
		return true
	}
	//delInernalNode(dstNode)
	delNode(dstNode)
	return true
}

func getSuccessor(node *bstNode) **bstNode {
	c := &node.right
	for ; (*c).left != nil; c = &((*c).left) {
	}
	return c
}

func delNode(node *bstNode) {
	successor := getSuccessor(node)
	fmt.Printf("node:%d successor:%d\n", node.key, (*successor).key)
	node.key = (*successor).key
	node.value = (*successor).value
	if (*successor).right == nil {
		*successor = nil
		return
	}
	r := (*successor).right
	(*successor).value = r.value
	(*successor).key = r.key
	(*successor).right = nil
}

func delRightEmpty(delNode *bstNode) {
	child := &delNode.left
	dstNode := &delNode.left
	for ; child != nil && *child != nil; child = &(*child).right {
		dstNode = child
	}
	targetNode := *dstNode
	*dstNode = nil
	if targetNode.left != nil {
		delNode.left = targetNode.left
	}
	fmt.Printf("dstNode:%v\n", targetNode)
	delNode.value = targetNode.value
	delNode.key = targetNode.key
}

func delInernalNode(node *bstNode) {
	if node.right == nil {
		delRightEmpty(node)
		return
	}
	child := &node.right
	dstNode := &node.right
	for ; child != nil && *child != nil; child = &(*child).left {
		dstNode = child
	}
	targetNode := *dstNode
	*dstNode = nil
	//TODO: targetNode.right is always nil
	if targetNode.right != nil {
		node.right = targetNode.right
	}
	node.value = targetNode.value
	node.key = targetNode.key
}

func (bst *BSTree) InOrder() []int {
	keys := make([]int, 0, bst.size)
	inOrder(bst.root, &keys)
	return keys
}

//TODO: why keys []int doesn't work
func inOrder(root *bstNode, keys *[]int) {
	if root == nil {
		return
	}
	inOrder(root.left, keys)
	//keys = append(keys, root.key)
	*keys = append(*keys, root.key)
	inOrder(root.right, keys)
}

func (bst *BSTree) BFSTraverse() {
	if bst.root == nil {
		return
	}
	queue := []*bstNode{bst.root}
	bstTraverse(queue)
}

func bstTraverse(queue []*bstNode) {
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
	bstTraverse(queue)
}
