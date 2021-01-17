package minheap

import "fmt"

//Value represents a value inside the MinHeap
type Value interface {
	greaterThan(Value) bool
}

//MinHeap struct has a slice that holds the heap values
type MinHeap struct {
	values []Value
}

//Insert adds an element to the heap
func (h *MinHeap) Insert(key Value) {
	h.values = append(h.values, key)
	h.minHeapifyUp(len(h.values) - 1)
}

//Extract returns the smallest key, and removes it from the heap
func (h *MinHeap) Extract() interface{} {
	extracted := h.values[0]

	l := len(h.values) - 1

	if l < 0 {
		fmt.Println("The heap is currently empty, cannot extract")
		return -1
	}

	h.values[0] = h.values[l]
	h.values = h.values[:l]

	h.minHeapifyDown(0)

	return extracted
}

//minHeapifyUp will heapify from bottom to top
func (h *MinHeap) minHeapifyUp(index int) {
	for h.values[parent(index)].greaterThan(h.values[index]) {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

//minHeapifyDown will heapify from top to bottom
func (h *MinHeap) minHeapifyDown(index int) {
	lastIndex := len(h.values) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex || h.values[r].greaterThan(h.values[l]) { // when left child is the only child or when left child is smaller
			childToCompare = l
		} else { // when right child is smaller
			childToCompare = r
		}

		if h.values[index].greaterThan(h.values[childToCompare]) {
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
func (h *MinHeap) swap(i1, i2 int) {
	h.values[i1], h.values[i2] = h.values[i2], h.values[i1]
}
