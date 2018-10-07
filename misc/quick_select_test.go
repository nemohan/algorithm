package misc

import (
	"testing"
)

func testCase(t *testing.T, selectFunc func([]int, int) int) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{2, 3, 4, 1, 5}
	c := []int{4, 1, 10, 8, 7, 12, 9, 2, 15}

	if s := selectFunc(a, 2); s != 2 {
		t.Fatalf("expect the second min:%d got:%d\n", 2, s)
	}
	if s := selectFunc(b, 2); s != 2 {
		t.Fatalf("expect the second min:%d got:%d\n", 2, s)
	}
	if s := selectFunc(c, 5); s != 8 {
		t.Fatalf("expect the second min:%d got:%d\n", 8, s)
	}
}

func TestQuickSelect(t *testing.T) {
	testCase(t, QuickSelect)
}

func TestQuickSelectRecursive(t *testing.T) {
	testCase(t, QuickSelectRecursive)
}
