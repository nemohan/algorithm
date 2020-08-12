package str

import (
	"testing"
)

var input = []string{
	"a",
	"awls",
	"sad",
	"sam",
	"same",
	"sap",
}

func createTrie(input []string) *Trie {
	trie := NewTrie()
	for _, str := range input {
		trie.Insert(str)
	}
	return trie
}
func TestFind(t *testing.T) {
	trie := createTrie(input)
	for _, str := range input {
		node := trie.Find(str)
		if node == nil || !node.end {
			t.Fatalf("want key:%s failed\n", str)
			return
		}
	}
}

func TestKeys(t *testing.T) {
	trie := createTrie(input)
	keys := trie.Keys()
	if len(keys) != len(input) {
		t.Fatalf("want:%v got:%v\n", input, keys)
		return
	}
	for i, key := range keys {
		if key != input[i] {
			t.Fatalf("want:%s got:%s\n", input[i], key)
			return
		}
	}
}
func TestKeysWithPrefix(t *testing.T) {
	trie := createTrie(input)
	want := map[string]bool{"sad": false, "sam": false, "same": false, "sap": false}
	keys := trie.KeysWithPrefix("sa")
	if len(keys) != len(want) {
		t.Fatalf("got:%v want:%#v\n", keys, want)
	}
	for _, key := range keys {
		if _, ok := want[key]; !ok {
			t.Fatalf("got:%s %v\n", key, []byte(key))
			return
		}
	}
}
