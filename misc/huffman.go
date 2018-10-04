package misc

import (
	"fmt"
)

type node struct {
	cha   rune
	code  uint32
	prob  probType
	left  *node
	right *node
}

type huffmanTree struct {
	root      *node
	table     []*node
	size      int
	codeTable map[rune]string
}

type codeword struct {
	size int
	bits string
}

type probType float32

func NewHuffmanTree(table map[rune]probType) map[rune]uint32 {
	tree := createHuffmanTree(table)
	tree.Inorder()
	fmt.Printf("code table:%v\n", tree.codeTable)
	codeTable := tree.genCode()
	return codeTable
}

func createHuffmanTree(table map[rune]probType) *huffmanTree {
	tree := &huffmanTree{
		table:     make([]*node, 0),
		size:      len(table),
		codeTable: make(map[rune]string),
	}
	for r, v := range table {
		n := &node{
			cha:  r,
			prob: v,
		}
		tree.table = append(tree.table, n)
	}
	tree.create()
	return tree
}

func (h *huffmanTree) create() {
	for h.size > 1 {
		left, right := h.chooseMinCha()
		parent := &node{prob: left.prob + right.prob,
			left:  left,
			right: right,
		}
		h.updateChaTable(parent)
	}
	h.root = h.table[0]
}

func (h *huffmanTree) chooseMinCha() (*node, *node) {
	for i := 0; i < h.size-1; i++ {
		max := i
		for j := i + 1; j < h.size; j++ {
			if h.table[max].prob < h.table[j].prob {
				max = j
			}
		}
		h.table[i], h.table[max] = h.table[max], h.table[i]
	}
	l, r := h.table[h.size-1], h.table[h.size-2]
	h.table[h.size-1] = nil
	h.table[h.size-2] = nil
	h.size -= 2
	if l.cha > r.cha {
		return r, l
	}
	return l, r
}

func (h *huffmanTree) updateChaTable(node *node) {
	h.table[h.size] = node
	h.size++
}

func (h *huffmanTree) genCode() map[rune]uint32 {
	return nil
}

func (h *huffmanTree) Inorder() {
	h.inorder(h.root, "")
}

func (h *huffmanTree) inorder(node *node, bits string) {
	if node == nil {
		return
	}
	h.inorder(node.left, bits+"0")
	fmt.Printf("prob:%v ch:%q bits:%s\n", node.prob, node.cha, bits)
	if node.left == nil && node.right == nil {
		h.codeTable[node.cha] = bits
	}

	if node.right != nil {
		bits += "1"
	}
	h.inorder(node.right, bits)
}
