package str

type tstNode struct {
	end    bool
	left   *tstNode
	middle *tstNode
	right  *tstNode
	ch     byte
}

type Tst struct {
	root *tstNode
	size int
}

func NewTst() *Tst {
	return &Tst{}
}

func newTstNode(ch byte, end bool) *tstNode {
	return &tstNode{
		ch:  ch,
		end: end,
	}
}

func (t *Tst) Find(key string) *tstNode {
	return tstFind(t.root, 0, key)
}

func tstFind(node *tstNode, index int, key string) *tstNode {
	if node == nil {
		return nil
	}
	k := key[index]
	if index == len(key)-1 {
		if node.ch == k && node.end {
			return node
		} else if node.ch == k && !node.end {
			return nil
		}
		//k less or greater than k

	}
	if node.ch == k {
		return tstFind(node.middle, index+1, key)
	} else if k < node.ch {
		return tstFind(node.left, index, key)
	}
	return tstFind(node.right, index, key)
}

func (t *Tst) Insert(key string) {
	t.root = tstInsert(t.root, 0, key)
}

/*
 a(true)  s
 w(true)  a
 l		  d(true)
 s(true)
*/
func tstInsert(node *tstNode, index int, key string) *tstNode {
	if node == nil {
		node = newTstNode(key[index], false)
		if index == len(key)-1 {
			node.end = true
			return node
		}
	}

	k := key[index]
	if node.ch == k {
		if index == len(key)-1 {
			node.end = true
			return node
		}
		node.middle = tstInsert(node.middle, index+1, key)
	} else if k < node.ch {
		node.left = tstInsert(node.left, index, key)
	} else {
		node.right = tstInsert(node.right, index, key)
	}
	return node
}

func (t *Tst) Keys() []string {
	keys := make([]string, 0)
	keys = tstCollect(t.root, "", keys)
	return keys
}

func tstCollect(node *tstNode, key string, keys []string) []string {
	if node.end {
		keys = append(keys, key+string(node.ch))
	}
	if node.middle != nil {
		keys = tstCollect(node.middle, key+string(node.ch), keys)
	}
	if node.left != nil {
		keys = tstCollect(node.left, key, keys)
	}
	if node.right != nil {
		keys = tstCollect(node.right, key, keys)
	}
	return keys
}

func (t *Tst) Delete(key string) {
	return
}

func tstDelete(node *tstNode) {

}
