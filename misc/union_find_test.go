package misc

import (
	"testing"
)

/*
func TestUnion(t *testing.T) {
	uset := MakeSet([]int{1, 2, 3, 4, 5})
	uset.Union(1, 2)
	elements := uset.Find(2)
	expected := map[int]bool{1: false, 2: false}
	for _, e := range elements {
		if _, ok := expected[e]; !ok {
			t.Fatalf("expected:%v got:%v\n", expected, elements)
		}
	}

}
*/

func TestUnion(t *testing.T) {
	uset := MakeSet(5)
	uset.Unite(1, 3)
	if !uset.IsSameSet(1, 3) {
		t.Fatalf("id 1 and 3 is not in the same set\n")
	}
	uset.Unite(0, 4)
	uset.Unite(0, 3)
	if !uset.IsSameSet(1, 0) {
		t.Fatalf("id 1 and 0 is not in the same set\n")
	}
}

func TestUnion2(t *testing.T) {
	uset := MakeSet(5)
	uset.Unite(3, 1)
	if !uset.IsSameSet(1, 3) {
		t.Fatalf("id 1 and 3 is not in the same set\n")
	}
	uset.Unite(4, 0)
	uset.Unite(3, 0)
	if !uset.IsSameSet(1, 0) {
		t.Fatalf("id 1 and 0 is not in the same set\n")
	}
}
