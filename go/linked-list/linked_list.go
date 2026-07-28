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

		if l.First() == nil && l.Last() == nil {
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
	if n == nil {
		return nil
	}

	return n.next
}

func (n *Node) Prev() *Node {
	if n == nil {
		return nil
	}

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
	if l.First() == nil && l.Last() == nil {
		nl := NewList(v)
		*l = *nl
	} else {
		newNode := Node{Value: v, previous: l.last}
		l.last.next = &newNode
		l.last = &newNode
	}
}

func (l *List) Shift() (value any, err error) {
	if l.First() == nil {
		err := errors.New("")
		return value, err
	}

	value = l.First().Value
	next := l.First().Next()
	l.first = next

	if l.First() != nil {
		next.previous = nil
	}

	if l.Count() == 0 {
		*l = *NewList()
	}

	return value, err
}

func (l *List) Pop() (value any, err error) {
	if l.Last() == nil {
		err := errors.New("")
		return value, err
	}

	if l.First() == l.Last() {
		value = l.First().Value
		l.first = nil
		l.last = nil

		return value, err
	}

	value = l.Last().Value
	l.last = l.Last().Prev()
	l.Last().next = nil

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
	node := lCopy.First()

	for node != nil {
		result = result + 1
		node = node.Next()
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
			if node.Prev() == nil {
				ll.first = node.Next()
			} else {
				node.Prev().next = node.Next()
				if node.Next() == nil {
					ll.last = node.Prev()
				} else {
					node.Next().previous = node.Prev()
				}
			}

			break
		}

		node = node.Next()
	}

	return found
}
