package search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	src := []int{1, 2, 3, 4, 5, 6}
	for _, e := range src {
		if !BinarySearch(src, e) {
			t.Fatalf("not find element:%d\n", e)
		}
	}
}

func TestLumotoPartion(t *testing.T) {
	old := []int{1, 2, 3, 4, 5, 6}
	src := []int{1, 2, 3, 4, 5, 6}
	lumotoPartition(src)
	for i, e := range old {
		if e != src[i] {
			t.Fatalf("expect:%v got:%v\n", old, src)
		}
	}

	fmt.Printf("%v\n", src)
	src1 := []int{3, 4, 1, 2, 6, 5}
	lumotoPartition(src1)
	if src1[2] != 3 {
		t.Fatalf("got:%v\n", src1)
	}
	fmt.Printf("%v\n", src1)
}
