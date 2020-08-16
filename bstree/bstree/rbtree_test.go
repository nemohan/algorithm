package bstree

import (
	"fmt"
	"testing"
)

func TestRBInsertOnly1Node(t *testing.T) {
	rb := NewRBTree()
	rb.Insert(10, nil)
	rb.Insert(2, nil)
	doFindTest(t, rb, []int{10, 2})
	checkRBRedLinkLeanLeft(t, rb)
	checkRBRedLinkNum(t, rb, 1)

	rb = NewRBTree()
	rb.Insert(10, nil)
	rb.Insert(20, nil)
	doFindTest(t, rb, []int{10, 20})
	checkRBRedLinkLeanLeft(t, rb)
	checkRBRedLinkNum(t, rb, 1)
}

func TestRBInsertOnly2Node(t *testing.T) {
	rb := NewRBTree()
	rb.Insert(10, nil)
	rb.Insert(2, nil)
	rb.Insert(1, nil)
	doFindTest(t, rb, []int{10, 2, 1})
	checkRBRedLinkLeanLeft(t, rb)
	checkRBRedLinkNum(t, rb, 0)

	rb = NewRBTree()
	rb.Insert(10, nil)
	rb.Insert(2, nil)
	rb.Insert(20, nil)
	doFindTest(t, rb, []int{10, 2, 20})
	checkRBRedLinkLeanLeft(t, rb)
	checkRBRedLinkNum(t, rb, 0)

	rb = NewRBTree()
	rb.Insert(10, nil)
	rb.Insert(2, nil)
	rb.Insert(8, nil)
	doFindTest(t, rb, []int{10, 2, 8})
	checkRBRedLinkLeanLeft(t, rb)
	checkRBRedLinkNum(t, rb, 0)
}

func TestRBInsertTo2Node(t *testing.T) {
	/*
		input := []int{19, 5, 1, 18, 3, 8 , 24}
		rb := NewRBTree()
		for _, v := range input {
			rb.Insert(v, nil)
		}
		dumpTree(rb.root)
		doFindTest(t, rb, input)
		checkRBRedLinkLeanLeft(t, rb)
		checkRBRedLinkNum(t, rb, 2)
	*/
	input := []int{19, 5, 1, 18, 3, 8, 24, 13, 16, 12}
	rb := NewRBTree()
	for _, v := range input {
		rb.Insert(v, nil)
	}
	dumpTree(rb.root)
	doFindTest(t, rb, input)
	checkRBRedLinkLeanLeft(t, rb)
	checkRBRedLinkNum(t, rb, 3)
}
func TestRBInsertTo3Node(t *testing.T) {

}

func TestRBInsert(t *testing.T) {
	input := genTestInput()
	rb := NewRBTree()
	for _, v := range input {
		rb.Insert(v, nil)
	}
	doFindTest(t, rb, input)
	checkRBRedLinkLeanLeft(t, rb)
}

func TestRBDelete(t *testing.T) {

}

//no node has two red links connected to it
func checkRBProperty1(t *testing.T) {

}

func checkRBBlackLinks(t *testing.T) {
}

func checkRBRedLinkLeanLeft(t *testing.T, rb *RBTree) {
	leanRight := false
	rbPreOrderTrav(rb.root, &leanRight, nil)
	if leanRight {
		t.Fail()
	}
}

func checkRBRedLinkNum(t *testing.T, rb *RBTree, wantNum int) {
	num := 0
	rbPreOrderTrav(rb.root, nil, &num)
	if wantNum != num {
		t.Errorf("want redlink num:%d got:%d\n", wantNum, num)
	}
}

func rbPreOrderTrav(node *RBNode, leanRight *bool, redLinkNum *int) {
	if node == nil {
		return
	}
	if leanRight != nil && isRed(node.right) {
		*leanRight = true
	}
	if redLinkNum != nil && isRed(node.left) {
		fmt.Printf("red node:%d\n", node.left.key)
		*redLinkNum++
	}
	rbPreOrderTrav(node.left, leanRight, redLinkNum)
	rbPreOrderTrav(node.right, leanRight, redLinkNum)
}
