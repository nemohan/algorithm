package bstree

import (
	"fmt"
	"reflect"
)

var debug = true

type Node interface {
	Left() Node
	Right() Node
	Key() int
}

func dumpTree(node Node) {
	dumpV2(node, 0)
}

func dumpV2(node Node, level int) {
	if node == nil {
		return
	}
	type nodeExt struct {
		Node
		level int
	}
	queue := []*nodeExt{&nodeExt{Node: node, level: level}}
	if debug {
		fmt.Printf("level:%d %s\n", queue[0].level, queue[0].Node)
	}
	for len(queue) > 0 {
		node := queue[0]
		left := node.Left()
		right := node.Right()
		lv := reflect.ValueOf(left)
		rv := reflect.ValueOf(right)
		level := node.level
		if left != nil && !lv.IsNil() {
			t := &nodeExt{Node: left, level: level + 1}
			queue = append(queue, t)
			if debug {
				fmt.Printf("left level:%d %s\n", t.level, left)
			}
		}
		if right != nil && !rv.IsNil() {
			t := &nodeExt{Node: right, level: level + 1}
			queue = append(queue, t)
			if debug {
				fmt.Printf("right level:%d %s\n", t.level, right)
			}
		}
		queue = queue[1:]
	}
}

func inorderTraverse(node Node, data []interface{}) []interface{} {
	v := reflect.ValueOf(node)
	if node == nil || v.IsNil() {
		return data
	}
	data = inorderTraverse(node.Left(), data)
	data = append(data, node.Key())
	data = inorderTraverse(node.Right(), data)
	return data
}

func dumpNodeHeight(node Node) int {
	v := reflect.ValueOf(node)
	if node == nil || v.IsNil() {
		return -1
	}
	lh := dumpNodeHeight(node.Left())
	rh := dumpNodeHeight(node.Right())
	h := lh
	if lh < rh {
		h = rh
	}
	h++
	if debug {
		fmt.Printf("node:%d height:%d\n", node.Key(), h)
	}
	return h
}

func isBalance(node Node) bool {
	balance := true
	defer func() {
		if r := recover(); r != nil {
			balance = false
		}
	}()
	return balance
}

func checkNodeHeight(node Node) int {
	if node == nil {
		return -1
	}
	lh := dumpNodeHeight(node.Left())
	rh := dumpNodeHeight(node.Right())
	h := lh
	if lh-rh > 1 || lh-rh < -1 {
		panic("no balance")
	}
	if lh < rh {
		h = rh
	}
	h++
	return h
}
