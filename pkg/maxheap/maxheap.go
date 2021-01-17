package maxheap

import "fmt"

//Value represents a value inside the MaxHeap
type Value interface {
	greaterThan(Value) bool
}

//MaxHeap struct has a slice that holds the heap values
type MaxHeap struct {
	values []Value
}

//Insert adds an element to the heap
func (h *MaxHeap) Insert(key Value) {
	h.values = append(h.values, key)
	h.maxHeapifyUp(len(h.values) - 1)
}

//Extract returns the largest key, and removes it from the heap
func (h *MaxHeap) Extract() interface{} {
	extracted := h.values[0]

	l := len(h.values) - 1

	if l < 0 {
		fmt.Println("The heap is currently empty, cannot extract")
		return -1
	}

	h.values[0] = h.values[l]
	h.values = h.values[:l]

	h.maxHeapifyDown(0)

	return extracted
}

//maxHeapifyUp will heapify from bottom to top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.values[index].greaterThan(h.values[parent(index)]) {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

//maxHeapifyDown will heapify from top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.values) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex || h.values[l].greaterThan(h.values[r]) { // when left child is the only child or when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}

		if h.values[childToCompare].greaterThan(h.values[index]) {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else { // found its right place
			return
		}
	}

}

//Returns the parent index
func parent(i int) int {
	return (i - 1) / 2
}

//Returns the left child index
func left(i int) int {
	return 2*i + 1
}

//Returns the right child index
func right(i int) int {
	return 2*i + 2
}

//Swaps keys in the values
func (h *MaxHeap) swap(i1, i2 int) {
	h.values[i1], h.values[i2] = h.values[i2], h.values[i1]
}
