package str

import (
	"testing"
)

func createRadixTree(keys []string) *RadixTree {
	tst := NewRadixTree()
	for _, str := range keys {
		tst.Insert(str)
	}
	return tst
}
func TestRadixFind(t *testing.T) {
	curInput := input
	tst := createRadixTree(curInput)
	for _, str := range curInput {
		node := tst.Find(str)
		if node == nil || !node.end {
			t.Fatalf("want key:%s failed %#v\n", str, node)
			return
		}
	}
}

func TestRadixKeys(t *testing.T) {
	curInput := input
	tst := createRadixTree(curInput)
	keys := tst.Keys()
	t.Logf("%v\n", keys)
}

func TestRadixDelete(t *testing.T) {
	for i, dc := range deleteCases {
		trie := createRadixTree(dc.input)
		trie.Delete(dc.key)
		keys := trie.Keys()
		if !checkKeys(t, dc.want, keys) {
			t.Fatalf("case:%d faile\n", i)
			return
		}
	}
}
