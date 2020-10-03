package stack

import (
	"sync"
)

// Item is an interface to store any data type in the Stack
type Item interface{}

// Stack is the struct that contains the items
type Stack struct {
	items []Item
	mutex sync.Mutex
}

// NewEmpty returns a new instance of Stack with no elements
func NewEmpty() *Stack {
	return &Stack{
		items: nil,
	}
}

// New returns a new instance of Stack with a list of specified elements
func New(items []Item) *Stack {
	return &Stack{
		items: items,
	}
}

// Push pushes an item to the top of the Stack instance
func (stack *Stack) Push(item Item) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	stack.items = append(stack.items, item)
}

// Pop removes the item on top of the Stack instance and returns the item
func (stack *Stack) Pop() Item {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	if len(stack.items) == 0 {
		return nil
	}

	lastItem := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]

	return lastItem
}

// Top returns the element on top of the Stack instance
func (stack *Stack) Top() Item {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	if len(stack.items) == 0 {
		return nil
	}

	lastItem := stack.items[len(stack.items)-1]

	return lastItem
}

// IsEmpty checks if the Stack instance is empty
func (stack *Stack) IsEmpty() bool {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	return len(stack.items) == 0
}

// Size returns the size of the Stack instance
func (stack *Stack) Size() int {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	return len(stack.items)
}

// Clear clears the items from the Stack instance
func (stack *Stack) Clear() {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	stack.items = nil
}
