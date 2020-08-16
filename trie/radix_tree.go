package trie

type radixTreeNode struct {
	key   string
	value interface{}
	end   bool
	//children []*radixTreeNode
	children map[byte]*radixTreeNode
}

const (
	keyEqual         = 0
	keyPrefix        = -1
	keyReversePrefix = -2
	keyNotEqual      = -3
)

type RadixTree struct {
	root *radixTreeNode
}

func NewRadixTree() *RadixTree {
	return &RadixTree{}
}

func newRadixTreeNode(key string, end bool) *radixTreeNode {
	return &radixTreeNode{
		key:      key,
		end:      end,
		children: make(map[byte]*radixTreeNode, 0),
	}
}
func (r *RadixTree) Insert(key string) {
	r.root = radixTreeInsert(r.root, key)
}

func match(src, dst string) (int, int) {
	srcLen := len(src)
	dstLen := len(dst)
	//root
	if srcLen == 0 {
		return keyNotEqual, 0
	}
	i := 0
	for ; i < srcLen && i < dstLen; i++ {
		if src[i] != dst[i] {
			return keyNotEqual, i
		}
	}
	//相等
	if i == srcLen && i == dstLen {
		return keyEqual, 0
	} else if i < srcLen && i == dstLen { //dst 是src 的前缀
		return keyPrefix, i
	}
	//src是dst的前缀
	return keyReversePrefix, i
}

//TODO: 如何快速的确定新节点该放入哪个分支, 一种方式是用哈希表即map,字符作为键，节点作为值
func radixTreeInsert(node *radixTreeNode, key string) *radixTreeNode {
	if node == nil {
		return newRadixTreeNode(key, true)
	}

	t, index := match(node.key, key)

	switch t {
	case keyEqual: //key ==  node.key
		node.end = true
	case keyNotEqual: //key != node.key
		if index == 0 && node.key != "" {
			root := newRadixTreeNode("", false)
			root.children[key[0]] = newRadixTreeNode(key, true)
			root.children[node.key[0]] = node
			return root
		} else if index == 0 {
			c := key[0]
			node.children[c] = radixTreeInsert(node.children[c], key[index:])
		} else {
			c := node.key[index]
			parentNode := newRadixTreeNode(node.key[:index], false)
			node.key = node.key[index:]
			parentNode.children[c] = node
			c = key[index]
			parentNode.children[c] = radixTreeInsert(parentNode.children[c], key[index:])
			return parentNode
		}
	case keyPrefix: // key is a prefix of node.key. so split current node
		c := node.key[index]
		pnode := newRadixTreeNode(node.key[:index], true)
		node.key = node.key[index:]
		pnode.children[c] = node
		return pnode
	case keyReversePrefix: // node.key is prefix of key
		c := key[index]
		node.children[c] = radixTreeInsert(node.children[c], key[index:])
	}
	/*
		if -1 == index { //key 是node.key的前缀, 分裂节点

		} else if 0 == index { //key 和node.key相等
			node.end = true
		} else if index > 0 {
			//键 部分字符相同
			c := key[index]
			node.children[c] = radixTreeInsert(node.children[c], key[index:])
		}
	*/
	return node
}

func (r *RadixTree) Find(key string) *radixTreeNode {
	return findRadixTreeNode(r.root, key)
}

func findRadixTreeNode(node *radixTreeNode, key string) *radixTreeNode {
	if node == nil {
		return nil
	}
	t, index := match(node.key, key)
	switch t {
	case keyEqual: //key ==  node.key
		if node.end {
			return node
		}
		return nil
	case keyNotEqual: //key != node.key
		c := key[index]
		return findRadixTreeNode(node.children[c], key[index:])
	case keyPrefix: // key is a prefix of node.key.
		return nil
	case keyReversePrefix: // node.key is prefix of key
		c := key[index]
		return findRadixTreeNode(node.children[c], key[index:])
	default:
		panic("unknown t ")
	}
}

func (r *RadixTree) Keys() []string {
	keys := make([]string, 0)
	return collectRadixTreeKeys(r.root, "", keys)
}
func collectRadixTreeKeys(node *radixTreeNode, key string, keys []string) []string {
	if node.end {
		keys = append(keys, key+node.key)
	}
	for _, c := range node.children {
		keys = collectRadixTreeKeys(c, key+node.key, keys)
	}
	return keys
}

func (r *RadixTree) Delete(key string) {
	r.root = deleteRadixTreeNode(r.root, key)
}

func deleteRadixTreeNode(node *radixTreeNode, key string) *radixTreeNode {
	if node == nil {
		return nil
	}
	t, index := match(node.key, key)
	switch t {
	case keyEqual: //key ==  node.key
		if node.end {
			if len(node.children) == 0 {
				node = nil
			} else {
				node.end = false
			}
			return node
		}
		return nil
	case keyNotEqual: //key != node.key
		c := key[index]
		child := deleteRadixTreeNode(node.children[c], key[index:])
		if child == nil {
			delete(node.children, c)
			//merge when node has only one child,
			//caution: may be we have grandchild
			if len(node.children) == 1 {
				for k, v := range node.children {
					//grandchild
					if len(v.children) != 0 {
						break
					}
					node.key = node.key + v.key
					node.end = true
					delete(node.children, k)
				}
			}
		} else {
			node.children[c] = child
		}

	case keyPrefix: // key is a prefix of node.key.
		return nil
	case keyReversePrefix: // node.key is prefix of key
		c := key[index]
		child := deleteRadixTreeNode(node.children[c], key[index:])

		if child == nil {
			delete(node.children, c)
			//merge when node has only one child
			if len(node.children) == 1 {
				for k, v := range node.children {
					if len(v.children) != 0 {
						break
					}
					node.key = node.key + v.key
					node.end = true
					delete(node.children, k)
				}
			}
		} else {
			node.children[c] = child
		}
	default:
		panic("unknown t ")
	}
	return node
}
