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
		newNode := Node{Value: element, previous: prevNode, next: nil}
		prevNode = &newNode

		if i == 0 {
			newList.first = &newNode
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
	newNode := Node{Value: v, previous: nil, next: l.first}
	l.first = &newNode
}

func (l *List) Push(v any) {
	// fmt.Println(l, v)
	if l.first == nil && l.last == nil {
		nl := NewList(v)
		*l = *nl
	} else {
		newNode := Node{Value: v, previous: l.last, next: nil}
		l.last.next = &newNode
		l.last = &newNode
	}
}

func (l *List) Shift() (any, error) {
	panic("Please implement the Shift function")
}

func (l *List) Pop() (any, error) {
	if l.last == nil {
		return -1, errors.New("")
	}

	value := l.last.Value
	newNode := Node{Value: l.last.previous.Value, next: nil, previous: nil}
	l.last = &newNode
	// fmt.Println(l.last, value)

	return value, nil
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
