package bstree

type Node interface {
	Left() Node
	Right() Node
	Key() int
}
