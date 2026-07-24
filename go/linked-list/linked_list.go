package linkedlist

import (
	"errors"
)

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.
type Node struct {
	Value    any
	next     *Node
	previous *Node
}

type List struct {
	first *Node
	last  *Node
}

func NewList(elements ...any) *List {
	length := len(elements)
	newList := List{}
	var prevNode *Node

	for i, element := range elements {
		newNode := Node{Value: element, previous: prevNode}
		prevNode = &newNode

		if i == 0 {
			newList.first = &newNode
			newList.last = &newNode
		}

		if i > 0 {
			prevNode.next = &newNode
		}

		if i == length-1 {
			newList.last = &newNode
			newNode.next = nil
		}
	}

	return &newList
}

func (n *Node) Next() *Node {
	return n.previous
}

func (n *Node) Prev() *Node {
	return n.next
}

func (l *List) Unshift(v any) {
	newNode := Node{Value: v, next: l.first}
	l.first = &newNode
}

func (l *List) Push(v any) {
	if l.first == nil && l.last == nil {
		nl := NewList(v)
		*l = *nl
	} else {
		newNode := Node{Value: v, previous: l.last}
		l.last.next = &newNode
		l.last = &newNode
	}
}

func (l *List) Shift() (value any, err error) {
	if l.first == nil {
		err := errors.New("")
		return value, err
	}

	value = l.first.Value
	next := l.first.next
	l.first = next

	return value, err
}

func (l *List) Pop() (value any, err error) {
	if l.first == nil {
		err := errors.New("")
		return value, err
	}

	if *l.first == *l.last {
		value = l.first.Value
		l.first = nil
		l.last = nil

		return value, err
	}

	value = l.last.Value
	l.last = l.first
	return value, err
}

func (l *List) Reverse() {
	panic("Please implement the Reverse function")
}

func (l *List) First() *Node {
	panic("Please implement the First function")
}

func (l *List) Last() *Node {
	panic("Please implement the Last function")
}

func (l *List) Count() int {
	panic("Please implement the Count function")
}

// Delete removes the first node in a list with a given value.
// Returns true if a node was removed.
func (ll *List) Delete(v any) bool {
	panic("Please implement the Delete function")
}
