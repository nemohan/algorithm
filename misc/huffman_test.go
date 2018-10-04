package misc

import (
	"fmt"
	"testing"
)

func TestHuffmanTree(t *testing.T) {
	table := map[rune]probType{'A': 0.35, 'B': 0.1, 'C': 0.2, 'D': 0.2, '_': 0.15}
	expectedCodes := map[rune]string{'A': "11", 'B': "100", 'C': "00", 'D': "01", '_': "101"}
	codeTable := NewHuffmanTree(table)
	for k, v := range codeTable {
		code, ok := expectedCodes[k]
		if !ok || code != v {
			t.Fatalf("expected code table:%v got:%v\n", expectedCodes, codeTable)
		}
	}
}

func TestEncode(t *testing.T) {
	//table := map[rune]probType{'A': 0.35, 'B': 0.1, 'C': 0.2, 'D': 0.2, '_': 0.15}
}

func TestDecode(t *testing.T) {
	table := map[rune]probType{'A': 0.4, 'B': 0.1, 'C': 0.2, 'D': 0.15, '_': 0.15}
	codeTable := NewHuffmanTree(table)
	decodeTable := make(map[string]rune)
	maxCodeLen := 0
	for k, v := range codeTable {
		if len(v) > maxCodeLen {
			maxCodeLen = len(v)
		}
		decodeTable[v] = k
	}
	decodedTxt := ""
	str := "100010111001010"
	for i := 0; i < len(str); {
		for j := i + 1; j <= maxCodeLen+i; j++ {
			prefix := str[i:j]
			v, ok := decodeTable[prefix]
			if ok {
				fmt.Printf("v:%v\n", v)
				i += j - i
				decodedTxt += string(v)
				break
			}
		}
	}
	dstStr := "ABADA_C"
	if decodedTxt != dstStr {
		t.Fatalf("expected decoded txt:%s got:%s\n", decodedTxt, dstStr)
	}
}
