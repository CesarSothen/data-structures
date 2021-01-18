package linkedlist

//SinglyLinkedList holds the head node of the linked list and its length
type SinglyLinkedList struct {
	head   *Node
	length int
}

//Get retrieves the Value v stored at index i
func (l *SinglyLinkedList) Get(i int) Value {
	return nil
}

//Add inserts a Value v at index i
func (l *SinglyLinkedList) Add(i int, v Value) bool {
	return false
}

//AddFirst inserts a Value v at the beggining of the LinkedList
func (l *SinglyLinkedList) AddFirst(v Value) {
	return
}

//AddLast inserts a Value v at the end of the LinkedList
func (l *SinglyLinkedList) AddLast(v Value) {
	return
}

//Delete removes and returns the first occurrence of a Value v from the LinkedList
func (l *SinglyLinkedList) Delete(v Value) Value {
	return nil
}

//Clear removes all Nodes from the LinkedList
func (l *SinglyLinkedList) Clear() {
	return
}

//Contains checks if the LinkedList contains a value v
func (l *SinglyLinkedList) Contains(v Value) bool {
	return false
}
