package tree

import (
	"fmt"
	"sync"
)

type btreeKey struct {
	key    int
	parent *btreeNode
	child  *btreeNode
	value  interface{}
	node   *btreeNode
}

type btreeNode struct {
	keys    []*btreeKey
	lastKey *btreeKey
	last    int
	parent  *btreeNode
	keyNum  int
}

type BTree struct {
	root   *btreeNode
	num    int
	min    int //min child number of a interal node
	max    int //max child number of a internal node
	rwLock *sync.RWMutex
	keyMax int
	keyMin int
}

func NewBTree(max int) *BTree {
	min := max / 2
	if max%2 != 0 {
		min += 1
	}
	root := newNode(max)
	return &BTree{
		root:   root,
		min:    min,
		max:    max,
		keyMax: max - 1,
		keyMin: min - 1,
		rwLock: new(sync.RWMutex),
	}

}

func newNode(max int) *btreeNode {
	return &btreeNode{
		keys:   make([]*btreeKey, max),
		keyNum: 0,
	}
}

func (pn *btreeNode) add(key int, value interface{}) (*btreeKey, *btreeKey) {
	i := pn.keyNum
	for ; i > 0; i-- {
		if pn.keys[i-1].key > key {
			pn.keys[i] = pn.keys[i-1]
			continue
		}
		break
	}
	neighbour := pn.lastKey
	keyNode := &btreeKey{key: key, value: value}
	pn.keys[i] = keyNode
	if i < pn.keyNum {
		neighbour = pn.keys[i+1]
	}
	if neighbour == nil {
		neighbour = &btreeKey{}
		pn.lastKey = neighbour
	}
	pn.keyNum++
	return keyNode, neighbour
}

func moveFromRight(start int, dstNode *btreeNode, srcNode *btreeNode) {
	for i := 0; i < start; i++ {
		j := i + start + 1
		dstNode.keys[i] = srcNode.keys[j]
		srcNode.keys[j] = nil
		dstNode.keyNum++
		srcNode.keyNum--
		//dstNode.lastKey = srcNode.keys[start]
	}
	srcNode.keyNum--
	//fix begin
	//the lastKey of dstNode must be empy
	//dstNode.lastKey = srcNode.lastKey.child
	dstNode.lastKey = &btreeKey{child: srcNode.lastKey.child}
	if srcNode.keys[start].child != nil {
		//TODO: check lastKey is nil
		srcNode.lastKey.child = srcNode.keys[start].child
	}
	srcNode.keys[start] = nil
	//fix begin
	/*
		srcNode.lastKey = nil
	*/
}

func (pn *btreeNode) dump() {
	fmt.Printf("begin----------->\n")
	for i := 0; i < pn.keyNum; i++ {
		fmt.Printf("key:%d %v\n", pn.keys[i].key, pn.keys[i])

	}
	fmt.Printf("last key:%v\n", pn.lastKey)
	fmt.Printf("end <-----------\n")
}

func (bt *BTree) split(pn *btreeNode) {
	//TODO: base condition
	if pn.keyNum <= bt.keyMax {
		return
	}

	left := pn
	medianIdx := bt.keyMin
	parent := pn.parent
	//root
	if parent == nil {
		parent = newNode(bt.max)
		bt.root = parent
	}
	//TODO: non root
	right := newNode(bt.max)
	left.parent = parent
	right.parent = parent

	midKey, neighbour := parent.add(left.keys[medianIdx].key, left.keys[medianIdx].value)
	midKey.child = left
	neighbour.child = right
	//fix begin
	/*
	 */
	//fix end
	parent.dump()
	moveFromRight(medianIdx, right, left)
	right.dump()
	bt.split(parent)
}

//insert sort
func (bt *BTree) add(dstNode *btreeNode, key int, value interface{}) {
	fmt.Printf("add key:%d keyNum:%d\n", key, dstNode.keyNum)
	dstNode.add(key, value)
	bt.split(dstNode)
}

func getMedian() {

}

//case 1: node number less than max
//case 2: node number equal max
//case 3:
func (bt *BTree) Insert(key int, value interface{}) {
	pn, ok := bt.find(bt.root.parent, bt.root, key)
	if ok {
		return
	}
	bt.add(pn, key, value)
}

func (bt *BTree) find(parent, pn *btreeNode, srcKey int) (*btreeNode, bool) {
	//leaf node
	if pn == nil {
		return parent, false
	}

	for i := 0; i < pn.keyNum; i++ {
		if pn.keys[i].key == srcKey {
			return pn, true
		}
		if pn.keys[i].key > srcKey {
			return bt.find(pn, pn.keys[i].child, srcKey)
		}
	}
	var last *btreeNode
	if pn.lastKey != nil {
		last = pn.lastKey.child
	}
	return bt.find(pn, last, srcKey)
}

func (bt *BTree) Del() {

}

func (bt *BTree) Dump() {
	queue := []*btreeNode{bt.root}
	i := 0
	traverse(i, queue)
}

func traverse(height int, queue []*btreeNode) {
	if len(queue) == 0 {
		return
	}
	num := len(queue)
	for i := 0; i < num; i++ {
		e := queue[0]
		queue = append(queue[:0], queue[1:]...)
		for k := 0; e != nil && k < e.keyNum; k++ {
			key := e.keys[k]
			fmt.Printf("layer i:%d key:%d parent:\n", height, key.key)
			queue = append(queue, key.child)
		}
		if e != nil && e.lastKey != nil {
			fmt.Printf("------->last layer i:%d key:%d\n\n", height, e.lastKey.key)
			queue = append(queue, e.lastKey.child)
		}
	}
	height++
	traverse(height, queue)
}
