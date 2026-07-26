package linkedlist

import (
	"errors"
	"fmt"
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
	l := &List{}
	var prevNode *Node

	for _, element := range elements {
		newNode := Node{Value: element, previous: prevNode}

		if l.first == nil && l.last == nil {
			l.first = &newNode
			l.last = &newNode
		} else {
			prevNode.next = &newNode
		}

		prevNode = &newNode
	}

	return l
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

	if l.Count() == 0 {
		*l = *NewList()
	}

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
	return l.first
}

func (l *List) Last() *Node {
	return l.last
}

func (l *List) Count() int {
	result := 0
	lCopy := *l
	node := lCopy.first

	for node != nil {
		result = result + 1
		node = node.next
	}

	return result
}

// Delete removes the first node in a list with a given value.
// Returns true if a node was removed.
func (ll *List) Delete(v any) bool {
	panic("Please implement the Delete function")
}
