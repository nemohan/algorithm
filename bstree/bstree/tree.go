package bstree

type Tree interface {
	Find(int) (interface{}, bool)
	Delete(int)
	Insert(int, interface{})
}
