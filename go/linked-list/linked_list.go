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
	l := &List{}
	var prevNode *Node

	for _, element := range elements {
		newNode := Node{Value: element, previous: prevNode}

		if l.first == nil && l.last == nil {
			l.first = &newNode
			l.last = &newNode
		} else {
			prevNode.next = &newNode
			l.last = &newNode
		}

		prevNode = &newNode
	}

	return l
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.previous
}

func (l *List) Unshift(v any) {
	current := l.First()

	newNode := &Node{Value: v, next: current}
	l.first = newNode

	if l.Count() < 2 {
		l.last = newNode
	} else {
		current.previous = newNode
	}
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
	l.last = l.last.previous
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
func (ll *List) Delete(v any) (found bool) {
	node := ll.First()

	for node != nil {
		if node.Value == v {
			found = true
			if ll.Count() == 1 {
				*ll = *NewList()
			}

			// node.previous == nil should mean it's the first element in the linked list
			if node.previous == nil {
				ll.first = node.next
			} else {
				node.previous.next = node.next
				if node.next == nil {
					ll.last = node.previous
				} else {
					node.next.previous = node.previous
				}
			}

			break
		}

		node = node.next
	}

	return found
}
