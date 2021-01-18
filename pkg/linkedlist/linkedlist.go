package linkedlist

//Value represents a value in a LinkedList Node
type Value interface{}

//Node holds its value and a pointer for its next node
type Node struct {
	value Value
	next  *Node
}

//LinkedList holds
type LinkedList interface {
	Get(i int) Value
	Add(i int, v Value) bool
	AddFirst(v Value)
	AddLast(v Value)
	Delete(v Value) Value
	Clear()
	Contains(v Value) bool
}
