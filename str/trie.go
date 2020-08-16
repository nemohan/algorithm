package str

import "fmt"

type Node struct {
	end   bool    //单次结束符
	next  []*Node //考虑哈希表或二叉查找树
	level int
}

type Trie struct {
	root *Node
	size int
}

func NewTrie() *Trie {
	return &Trie{
		size: 0,
		root: newNode(),
	}
}

func newNode() *Node {
	return &Node{
		end:  false,
		next: make([]*Node, 128),
	}
}

func (t *Trie) Insert(key string) {
	node := t.Find(key)
	if node != nil {
		return
	}
	insert(t.root, 0, key)
}
func insert(node *Node, index int, key string) {
	i := int(key[index])
	if node.next[i] == nil {
		node.next[i] = newNode()
	}
	if index == len(key)-1 {
		node.next[i].end = true
		return
	}
	insert(node.next[i], index+1, key)
}

func (t *Trie) Find(key string) *Node {
	return find(t.root, 0, key)
}

func find(node *Node, index int, key string) *Node {
	i := int(key[index])
	next := node.next[i]
	if next == nil {
		return nil
	} else if index == len(key)-1 && !next.end {
		return nil
	} else if index == len(key)-1 && next.end {
		return next
	}
	return find(next, index+1, key)
}

func (t *Trie) Delete(key string) {
	deleteKeyInTrie(t.root, 0, key)
}

func hasChildren(node *Node) bool {
	for _, p := range node.next {
		if p != nil {
			return true
		}
	}
	return false
}
func deleteKeyInTrie(node *Node, index int, key string) *Node {
	k := key[index]
	next := node.next[k]
	if next == nil {
		return node
	}
	if index == len(key)-1 {
		if next.end {
			if hasChildren(next) {
				next.end = false
			} else {
				node.next[k] = nil
			}
			return node
		} else {
			return node
		}
	}
	node.next[k] = deleteKeyInTrie(next, index+1, key)
	if hasChildren(node.next[k]) {
		return node
	}
	if !node.next[k].end {
		node.next[k] = nil
	}
	return node
}

func (t *Trie) Dump() {
	t.root.level = 1
	queue := []*Node{t.root}
	for len(queue) > 0 {
		node := queue[0]
		fmt.Printf("level:%d", node.level)
		for i, p := range node.next {
			if p != nil {
				fmt.Printf(" %s-%v ", string(i), p.end)
				p.level = node.level + 1
				queue = append(queue, p)
			}
		}
		fmt.Printf("\n")
		queue = queue[1:]
	}
}

func (t *Trie) Keys() []string {
	keys := make([]string, 0)
	for i, n := range t.root.next {
		if n == nil {
			continue
		}
		s := string(i)
		if n.end {
			keys = append(keys, s)
		}
		keys = collectV2(string(i), n, keys)
	}
	return keys
}

func collectV2(key string, node *Node, keys []string) []string {
	for i, n := range node.next {
		if n == nil {
			continue
		}
		//key = key + string(i)
		if n.end {
			keys = append(keys, key+string(i))
		}
		keys = collectV2(key+string(i), n, keys)
	}
	return keys
}

func (t *Trie) KeysWithPrefix(key string) []string {
	return collect(t.root, 0, key)
}

func collect(node *Node, index int, key string) []string {
	k := key[index]
	if node.next[k] == nil {
		return nil
	}
	if index == len(key)-1 {
		keys := make([]string, 0)
		if node.next[k].end {
			keys = append(keys, key)
		}
		for _, pnext := range node.next {
			if pnext == nil {
				continue
			}
			//keys = collectV2(key+string(i), pnext, keys)
			keys = collectV2(key, pnext, keys)
		}
		return keys
	}
	return collect(node.next[k], index+1, key)
}
