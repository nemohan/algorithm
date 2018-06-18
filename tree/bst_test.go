package tree

import (
	//"fmt"
	"testing"
)

var srcData = []int{20, 15, 21, 10, 18, 22, 5, 14, 16, 19}
var srcDataSorted = []int{5, 10, 14, 15, 16, 18, 19, 20, 21, 22}
var srcDataDelLeaf = []int{5, 10, 15, 16, 18, 19, 20, 21, 22}
var srcDataDelRightEmpty = []int{20, 15, 10, 18, 5, 14, 16, 19}

func createTree() *BSTree {
	bst := NewBSTree()
	for i, key := range srcData {
		bst.Insert(key, i)
	}
	return bst
}
func TestInsert(t *testing.T) {
	bst := createTree()
	_, ok := bst.Find(10)
	if !ok {
		t.Fatalf("not find key:%d in tree\n", 10)
	}
}

func TestInOrder(t *testing.T) {
	bst := createTree()
	keys := bst.InOrder()
	checkResult(keys, srcDataSorted, t)
}

func checkResult(result []int, expected []int, t *testing.T) {
	keys := result
	//fmt.Printf("keys:%v\n", keys)
	if len(keys) != len(expected) {
		t.Fatalf("expected keys number:%d got %d\n", len(expected), len(keys))
	}
	for i, v := range expected {
		if v != keys[i] {
			t.Fatalf("expected keys:%v got:%v\n", expected, keys)
		}
	}
}
func TestDelLeaf(t *testing.T) {
	bst := createTree()
	bst.Del(14)
	keys := bst.InOrder()
	checkResult(keys, srcDataDelLeaf, t)
}

func checkDecl(bst *BSTree, key int, t *testing.T) {
	bst.Del(key)
	keys := bst.InOrder()
	expected := make([]int, 0, len(srcDataSorted)-1)
	for _, v := range srcDataSorted {
		if v == key {
			continue
		}
		expected = append(expected, v)
	}
	checkResult(keys, expected, t)
}

func TestDelHasTwoSubTree(t *testing.T) {
	bst := createTree()
	checkDecl(bst, 20, t)

	bst = createTree()
	checkDecl(bst, 10, t)
	bst = createTree()
	checkDecl(bst, 15, t)
}

func TestDelInternal(t *testing.T) {
	bst := createTree()
	bst.Del(22)
	bst.InOrder()

	bst.Del(21)
	bst.InOrder()

	bst.Del(20)
	bst.InOrder()
}
