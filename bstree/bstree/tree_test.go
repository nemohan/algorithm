package bstree

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

func doFindTest(t *testing.T, tree Tree, input []int) {
	for _, v := range input {
		if _, ok := tree.Find(int(v)); !ok {
			//dumpTree(tree.root)
			t.Errorf("input:%v key:%d\n", input, v)
			return
		}
	}
}

func genTestInputV2(inputSize, size int) []int {
	keySet := make(map[int]bool, 0)
	keys := make([]int, 0)
	for i := 0; i < inputSize; i++ {
		//j := rand.Int31n(500)
		j := rand.Intn(size)
		if _, ok := keySet[j]; ok {
			continue
		}
		keySet[j] = false
		keys = append(keys, j)
	}
	return keys

}
func genTestInput() []int {
	keySet := make(map[int]bool, 0)
	keys := make([]int, 0)
	for i := 0; i < 200; i++ {
		//j := rand.Int31n(500)
		j := rand.Intn(500)
		if _, ok := keySet[j]; ok {
			continue
		}
		keySet[j] = false
		keys = append(keys, j)
	}
	return keys
}

func genInputFromStr(str string) []int {
	items := strings.Split(str, " ")
	input = make([]int, len(items))
	i := 0
	for _, v := range items {
		input[i], _ = strconv.Atoi(v)
		i++
	}
	return input
}
func TestIsBalance(t *testing.T) {
	input := make([]int, 0)
	for i := 1; i < 20; i++ {
		input = append(input, i)
	}
	tree := newTestBSTree(input)
	if !isBalance(tree.root) {
		t.Fail()
	}
}
