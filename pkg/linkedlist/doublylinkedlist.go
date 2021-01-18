package linkedlist

//DoublyLinkedList holds the head node of the linked list, the tail and its length
type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

//Get retrieves the Value v stored at index i
func (l *DoublyLinkedList) Get(i int) Value {
	return nil
}

//Add inserts a Value v at index i
func (l *DoublyLinkedList) Add(i int, v Value) bool {
	return false
}

//AddFirst inserts a Value v at the beggining of the LinkedList
func (l *DoublyLinkedList) AddFirst(v Value) {
	return
}

//AddLast inserts a Value v at the end of the LinkedList
func (l *DoublyLinkedList) AddLast(v Value) {
	return
}

//Delete removes and returns the first occurrence of a Value v from the LinkedList
func (l *DoublyLinkedList) Delete(v Value) Value {
	return nil
}

//Clear removes all Nodes from the LinkedList
func (l *DoublyLinkedList) Clear() {
	return
}

//Contains checks if the LinkedList contains a value v
func (l *DoublyLinkedList) Contains(v Value) bool {
	return false
}
