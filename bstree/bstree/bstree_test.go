package bstree

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

var data = []int{1, 4, 5, 2, 3}

//this input will cause TestInsert2 fail
var rawInput = `422 45 494 306 499 24 420 44 259 7 124 357 353 274 475 240 89 314 339 81 131 51 485 129 203 195 92 372 275 377 262 473 17 307 280 205 406 474 223 394 33 74 110 30 395 414 23 150 409 277 145 171 208 198 234 379 354 221 87 326 10 421 267 301 349 488 374 401 140 96 466 75 413 214 459 325 446 244 63 436 487 425 26 219 200 114 230 176 300 434 119 166 41 142 239 194 201 250 88 364 60 216 82 316 313 16 322 55 463 217 258 232 179 375 333 127 34 52 428 37 147 469 212 346 9 104 356 442 264 453 360 213 328 257 184 283 350 315 40 492 342 91 398 207 412 452 255 269 98 90 417 382 370 182 268 476 196`
var input []int

func init() {
	items := strings.Split(rawInput, " ")
	input = make([]int, len(items))
	i := 0
	for _, v := range items {
		input[i], _ = strconv.Atoi(v)
		i++
	}
}
func newTestBSTree(input []int) *BSTree {
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

func TestBSTInsert(t *testing.T) {
	tree := NewBSTree()
	for _, v := range data {
		tree.Insert(v, nil)
	}
	tree.InOrderTraverse()
}

func TestBSTInsert2(t *testing.T) {
	input = genTestInput()
	tree := NewBSTree()
	for _, v := range input {
		tree.Insert(int(v), nil)
	}
	doFindTest(t, tree, input)
}

func TestBSTDelLeaf(t *testing.T) {
	input := []int{20, 30, 25, 10, 5, 15, 12, 14, 13}
	expect := []int{5, 10, 12, 13, 14, 15, 20, 25, 30}
	tree := newTestBSTree(input)

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
	tree := newTestBSTree(input)

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
	tree := newTestBSTree(input)

	tree.Delete(10)
	output := make([]interface{}, 0)
	output = inorderTraverse(tree.root, output)
	expect = gotExpect(expect, 10)
	checkResult("TestBSTDelNodeWithTwoChild", t, expect, output)

	input = []int{20, 30, 25, 10, 5, 15}
	expect = []int{5, 10, 15, 20, 25, 30}
	tree = newTestBSTree(input)

	tree.Delete(10)
	output = make([]interface{}, 0)
	output = inorderTraverse(tree.root, output)
	expect = gotExpect(expect, 10)
	checkResult("TestBSTDelNodeWithTwoChild", t, expect, output)
	dumpTree(tree.root)
}

func TestBSTDelete(t *testing.T) {
	input = genTestInput()
	tree := newTestBSTree(input)
	num := len(input) / 2
	doFindTest(t, tree, input)
	for i := 0; i < num; i++ {
		k := input[i]
		tree.Delete(k)
		if _, ok := tree.Find(k); ok {
			t.Errorf("input:%v k:%d\n", input, k)
			return
		}
	}

}

func BenchmarkBSTFind(b *testing.B) {
	input := genTestInputV2(100000, 1000000)
	tree := newTestBSTree(input)
	for _, v := range input {
		tree.Insert(v, nil)
	}
	size := len(input)
	for i := 0; i < b.N; i++ {
		r := rand.Intn(1000000) % size
		if _, ok := tree.Find(input[r]); !ok {
			//b.Error("failed")
			b.Fail()
			return
		}
	}
}
