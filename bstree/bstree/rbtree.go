package bstree

import "fmt"

/*
	implementation of left-leaning red-black tree. which 1-1 coresponds with 2-3 tree
*/
type RBNode struct {
	key   int
	value interface{}
	color NodeColoer
	left  *RBNode
	right *RBNode
}

type NodeColoer bool

const (
	NodeColoerRed   = false
	NodeColoerBlack = true
)

type RBTree struct {
	root *RBNode
}

func NewRBTree() *RBTree {
	return &RBTree{}
}
func (rb *RBTree) Find(key int) (interface{}, bool) {
	node := rbFind(rb.root, key)
	if node == nil {
		return nil, false
	}
	return node.value, true
}

func rbFind(node *RBNode, key int) *RBNode {
	if node == nil {
		return nil
	}
	if key < node.key {
		return rbFind(node.left, key)
	} else if key > node.key {
		return rbFind(node.right, key)
	}
	return node
}

func (rb *RBTree) Insert(key int, value interface{}) {
	rb.root = rbInsert(rb.root, key, value)
	paintBlack(rb.root)
}

func rbInsert(node *RBNode, key int, value interface{}) *RBNode {
	if node == nil {
		return &RBNode{key: key, value: value, color: NodeColoerRed}
	}
	if key < node.key {
		node.left = rbInsert(node.left, key, value)
	} else if key > node.key {
		node.right = rbInsert(node.right, key, value)
	}
	//fmt.Printf("right:%v left:%v node:%d \n", node.left, node.right, node.key)
	if isRed(node.right) && isRed(node.left) {
		//fmt.Printf("flip node:%d \n", node.key)
		flipColors(node)
	} else if isRed(node.right) {
		node = rbLeftRotate(node)
	} else if node.left != nil && isRed(node.left) && isRed(node.left.left) {
		node = rbRightRotate(node)
		flipColors(node)
	}
	return node
}

func (rb *RBTree) Delete(key int) {

}
func rbDelete(node *RBNode, key int) *RBNode {
	if node == nil {
		return nil
	}
	if key < node.key {
		node.left = rbDelete(node.left, key)
	} else if key > node.key {
		node.right = rbDelete(node.right, key)
	}
	if node.left == nil && node.right == nil {
		return nil
	} else if node.left == nil && node.right != nil {
		return node.right
	} else if node.right == nil && node.left != nil {
		return node.left
	}
	//two sub trees
	minNode := findRBNodeMin(node.right)
	node.key = minNode.key
	node.value = minNode.value
	node.right = rbDelete(node.right, node.key)
	if isRed(node.right) && isRed(node.left) {
		//fmt.Printf("flip node:%d \n", node.key)
		flipColors(node)
	} else if isRed(node.right) {
		node = rbLeftRotate(node)
	} else if node.left != nil && isRed(node.left) && isRed(node.left.left) {
		node = rbRightRotate(node)
		flipColors(node)
	}
	return node
}

func findRBNodeMin(node *RBNode) *RBNode {
	if node.left == nil {
		return node
	}
	return findRBNodeMin(node.left)
}
func isRed(node *RBNode) bool {
	if node == nil {
		return false
	}
	return node.color == NodeColoerRed
}

func rbRightRotate(node *RBNode) *RBNode {
	lchild := node.left
	r := lchild.right
	lchild.right = node
	node.left = r
	paintColor(node.color, lchild)
	paintRed(node)
	//do we need change color of node and lchild ??
	return lchild
}

func rbLeftRotate(node *RBNode) *RBNode {
	rchild := node.right
	l := rchild.left
	rchild.left = node
	node.right = l
	paintColor(node.color, rchild)
	paintRed(node)
	return rchild
}

func paintColor(color NodeColoer, node *RBNode) {
	if node == nil {
		return
	}
	node.color = color
}
func paintBlack(node *RBNode) {
	if node == nil {
		panic("paint black nil node")
	}
	node.color = NodeColoerBlack
}
func paintRed(node *RBNode) {
	if node == nil {
		return
	}
	node.color = NodeColoerRed
}

func flipColors(node *RBNode) {
	paintBlack(node.right)
	paintBlack(node.left)
	paintRed(node)
}

func (rb *RBNode) Left() Node {
	return rb.left
}

func (rb *RBNode) Right() Node {
	return rb.right
}

func (rb *RBNode) Key() int {
	return rb.key
}

func (rb *RBNode) String() string {
	color := "black"
	if rb.color == NodeColoerRed {
		color = "red"
	}
	return fmt.Sprintf("{key:%d color:%s}\n", rb.key, color)
}
