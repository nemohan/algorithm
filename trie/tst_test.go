package str

import (
	"sort"
	"testing"
)

var tstInput = []string{
	"a",
	"awls",
	"aw",
	"sad",
	"sam",
	"same",
	"sap",
	"sa",
}

func createTst(keys []string) *Tst {
	tst := NewTst()
	for _, str := range keys {
		tst.Insert(str)
	}
	return tst
}
func TestTstFind(t *testing.T) {
	curInput := input
	tst := createTst(curInput)
	for _, str := range curInput {
		node := tst.Find(str)
		if node == nil || !node.end {
			t.Fatalf("want key:%s failed %#v\n", str, node)
			return
		}
	}
}

func TestTstKeys(t *testing.T) {
	trie := createTst(input)
	keys := trie.Keys()

	sort.Strings(keys)
	if len(keys) != len(sortInput) {
		t.Fatalf("want:%v got:%v\n", sortInput, keys)
		return
	}
	for i, key := range keys {
		t.Logf("key:%s\n", key)
		if key != sortInput[i] {
			t.Fatalf("want:%s got:%s\n", sortInput[i], key)
			return
		}
	}
}

func TestTstDelete(t *testing.T) {
	for _, dc := range deleteCases {
		trie := createTst(dc.input)
		trie.Delete(dc.key)
		keys := trie.Keys()
		checkKeys(t, dc.want, keys)
	}
}
