package trie

import (
	"sort"
	"testing"
)

type testCase struct {
	input []string
	want  []string
	key   string
}

var (
	input = []string{
		"a",
		"awls",
		"aw",
		"sad",
		"sam",
		"same",
		"sap",
		"sa",
	}
	deleteCases = []testCase{
		{
			key:   "same",
			input: input,
			want: []string{
				"a",
				"awls",
				"sad",
				"sam",
				"sap",
				"sa",
			},
		},

		{
			key:   "sam",
			input: input,
			want: []string{
				"a",
				"awls",
				"sad",
				"same",
				"sap",
				"sa",
			},
		},
		{
			key:   "sap",
			input: input,
			want: []string{
				"a",
				"awls",
				"sad",
				"sam",
				"same",
				"sa",
			},
		},
		{
			key:   "ax",
			input: input,
			want:  input,
		},
	}
	sortInput = []string{}
)

func init() {
	for i, dc := range deleteCases {
		want := make([]string, 0)
		for _, key := range dc.input {
			if key == dc.key {
				continue
			}
			want = append(want, key)
		}
		deleteCases[i].want = want
	}
	sortInput = make([]string, len(input))
	copy(sortInput, input)
	sort.Strings(sortInput)
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
	sort.Strings(keys)
	for i, key := range keys {
		if key != sortInput[i] {
			t.Fatalf("want:%s got:%s %v %d\n", sortInput[i], key, keys, i)
			return
		}
	}
}
func TestKeysWithPrefix(t *testing.T) {
	trie := createTrie(input)
	want := map[string]bool{"sa": false, "sad": false, "sam": false, "same": false, "sap": false}
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

func TestDelete(t *testing.T) {
	for i, dc := range deleteCases {
		trie := createTrie(dc.input)
		trie.Delete(dc.key)
		keys := trie.Keys()
		if !checkKeys(t, dc.want, keys) {
			t.Fatalf("case:%d faile\n", i)
			return
		}
	}
}

func checkKeys(t *testing.T, want, got []string) bool {
	sort.Strings(want)
	sort.Strings(got)
	if len(got) != len(want) {
		t.Logf("want:%v got:%v\n", want, got)
	}
	for i, w := range want {
		if w != got[i] {
			t.Logf("want:%s got:%s\n", w, got[i])
			return false
		}
	}
	return true
}

func TestInput(t *testing.T) {
	t.Logf("%v\n", input)
}
