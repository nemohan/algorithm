package str

import "fmt"

//Search search pattern in txt, return first index of patern
func Search(txt string, pattern string) int {
	if len(pattern) > len(txt) {
		return -1
	}
	return horspoolSearch(txt, pattern)
}

func constructShiftTable(pattern string) []int {
	table := make([]int, 128)
	n := len(pattern)
	for i := range table {
		table[i] = n
	}
	for i := 0; i < n-1; i++ {
		table[pattern[i]] = n - i - 1
	}
	for i := int("A"[0]); i <= int("Z"[0]); i++ {
		fmt.Printf("%s %d\n", string(i), table[i])
	}
	return table
}
func horspoolSearch(txt string, pattern string) int {
	table := constructShiftTable(pattern)
	n := len(pattern)
	txtLen := len(txt)
	for i := 0; i <= txtLen-n; {
		j := n - 1
		for j >= 0 && txt[i+j] == pattern[j] {
			j--
		}
		if j == -1 {
			return i
		}
		k := txt[i+n-1]
		i += table[k]
	}
	return -1
}

func bruteForceSearch(txt, pattern string) int {
	return -1
}

func bmSearch(txt string, pattern string) int {
	return -1
}
