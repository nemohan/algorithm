package bstree

import (
	"testing"
)

var data = []int{1, 4, 5, 2, 3}

func TestFind(t *testing.T) {

}

func TestInsert(t *testing.T) {
	tree := NewBSTree()
	for _, v := range data {
		tree.Insert(v, nil)
	}
	tree.InOrderTraverse()
}

func TestDump(t *testing.T) {
	tree := NewBSTree()
	for _, v := range data {
		tree.Insert(v, nil)
	}
}

func newTree(input []int) *BSTree {
	tree := NewBSTree()
	for _, v := range input {
		tree.Insert(v, nil)
	}
	return tree
}

func checkResult(f string, t *testing.T, input []int, output []interface{}) {
	for i, v := range input {
		o := output[i].(int)
		if v != o {
			t.Errorf("%s expect:%v got:%v\n", f, input, output)
			return
		}
	}
}
func gotExpect(input []int, k int) []int {
	dst := make([]int, len(input)-1)
	i := 0
	for _, v := range input {
		if v == k {
			continue
		}
		dst[i] = v
		i++
	}
	return dst
}
func TestBSTDelLeaf(t *testing.T) {
	input := []int{20, 30, 25, 10, 5, 15, 12, 14, 13}
	expect := []int{5, 10, 12, 13, 14, 15, 20, 25, 30}
	tree := newTree(input)

	tree.Delete(25)
	output := make([]interface{}, 0)
	output = inorderTraverse(tree.root, output)
	expect = gotExpect(expect, 25)
	checkResult("TestDelLeaf", t, expect, output)

	tree.Delete(5)
	output = make([]interface{}, 0)
	output = inorderTraverse(tree.root, output)
	checkResult("TestBSTDelLeaf", t, gotExpect(expect, 5), output)
}

func TestBSTDelNodeWithOneChild(t *testing.T) {
	input := []int{20, 30, 25, 10, 5, 15, 12, 14, 13}
	expect := []int{5, 10, 12, 13, 14, 15, 20, 25, 30}
	tree := newTree(input)

	tree.Delete(30)
	output := make([]interface{}, 0)
	output = inorderTraverse(tree.root, output)
	expect = gotExpect(expect, 30)
	checkResult("TestBSTDelNodeWithOneChild", t, expect, output)

	tree.Delete(12)
	output = make([]interface{}, 0)
	output = inorderTraverse(tree.root, output)
	checkResult("TestBSTDelNodeWithOneChild", t, gotExpect(expect, 12), output)
}

func TestBSTDelNodeWithTwoChild(t *testing.T) {
	input := []int{20, 30, 25, 10, 5, 15, 12, 14, 13}
	expect := []int{5, 10, 12, 13, 14, 15, 20, 25, 30}
	tree := newTree(input)

	tree.Delete(10)
	output := make([]interface{}, 0)
	output = inorderTraverse(tree.root, output)
	expect = gotExpect(expect, 10)
	checkResult("TestBSTDelNodeWithTwoChild", t, expect, output)

	input = []int{20, 30, 25, 10, 5, 15}
	expect = []int{5, 10, 15, 20, 25, 30}
	tree = newTree(input)

	tree.Delete(10)
	output = make([]interface{}, 0)
	output = inorderTraverse(tree.root, output)
	expect = gotExpect(expect, 10)
	checkResult("TestBSTDelNodeWithTwoChild", t, expect, output)
	dumpTree(tree.root)
}

func BenchmarkFind(t *testing.B) {

}
