package trie

import (
	"sort"
	"testing"
)

var tstInput = []string{
	"app",
	"apple",
	"beer",
	"add",
	"jam",
	"rental",
}
var tstSortInput = []string{
	"add",
	"app",
	"apple",
	"beer",
	"jam",
	"rental",
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
		t.Logf("%s\n", key)
		if key != sortInput[i] {
			t.Fatalf("want:%s got:%s\n", sortInput[i], key)
			return
		}
	}
}

func TestTstKeysWithPrefix(t *testing.T) {
	for i, pc := range prefixCases {
		tst := createTst(pc.input)
		keys := tst.KeysWithPrefix(pc.key)
		if !checkKeys(t, pc.want, keys) {
			t.Fatalf("case:%d want:%v got:%v\n", i, pc.want, keys)
			return
		}
	}
	tst := createTst(tstInput)
	want := []string{"add"}
	keys := tst.KeysWithPrefix("ad")
	if !checkKeys(t, want, keys) {
		t.Fatalf(" want:%v got:%v\n", want, keys)
		return
	}
}

/*
func TestTstDelete(t *testing.T) {
	for _, dc := range deleteCases {
		trie := createTst(dc.input)
		trie.Delete(dc.key)
		keys := trie.Keys()
		checkKeys(t, dc.want, keys)
	}
}
*/
