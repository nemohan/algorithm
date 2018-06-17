package tree

import (
	"fmt"
	"testing"
)

func TestNodeAdd(t *testing.T) {
	pn := newNode(5)
	pn.add(1, 1)
	pn.add(4, 4)
	pn.add(2, 2)
	pn.add(0, 0)
	if pn.keys[2].key != 2 {
		t.Fatal("wrong key order\n")
	}
	for i := 0; i < pn.keyNum; i++ {
		k := pn.keys[i]
		fmt.Printf("key:%d\n", k.key)
	}
}

func TestDump(t *testing.T) {
	bt := NewBTree(5)
	bt.Insert(100, 100)
	bt.Insert(110, 110)
	bt.Insert(120, 120)
	bt.Insert(130, 130)
	bt.Insert(140, 140)
	bt.Insert(150, 150)
	bt.Insert(160, 160)
	bt.Insert(170, 170)
	bt.Dump()

}
