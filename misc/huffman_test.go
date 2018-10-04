package misc

import (
	"testing"
)

func TestHuffmanTree(t *testing.T) {
	table := map[rune]probType{'A': 0.35, 'B': 0.1, 'C': 0.2, 'D': 0.2, '_': 0.15}
	//table := map[rune]int{'A': 35, 'B': 10, 'C': 20, 'D': 20, '_': 15}
	NewHuffmanTree(table)
}
