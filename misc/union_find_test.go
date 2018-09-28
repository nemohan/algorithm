package misc

import (
	"testing"
)

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
