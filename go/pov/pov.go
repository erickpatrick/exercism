package pov

type Tree struct {
	root     string
	children []*Tree
}

// New creates and returns a new Tree with the given root value and children.
func New(value string, children ...*Tree) *Tree {
	return &Tree{root: value, children: children}
}

// Value returns the value at the root of a tree.
func (tr *Tree) Value() string {
	return tr.root
}

// Children returns a slice containing the children of a tree.
// There is no need to sort the elements in the result slice,
// they can be in any order.
func (tr *Tree) Children() []*Tree {
	return tr.children
}

// String describes a tree in a compact S-expression format.
// This helps to make test outputs more readable.
// Feel free to adapt this method as you see fit.
func (tr *Tree) String() string {
	if tr == nil {
		return "nil"
	}
	result := tr.Value()
	if len(tr.Children()) == 0 {
		return result
	}
	for _, ch := range tr.Children() {
		result += " " + ch.String()
	}
	return "(" + result + ")"
}

// POV problem-specific functions

// FromPov returns the pov from the node specified in the argument.
func (tr *Tree) FromPov(from string) *Tree {
	if tr.Value() == from {
		return tr
	}

	var root *Tree
	var found *Tree
	parent := New(tr.Value())

	for _, child := range tr.Children() {
		found = child.FromPov(from)

		// fmt.Println("found", found, "| parent", parent)

		if found == nil {
			parent.children = append(parent.children, child)
			continue
		} else {
			var children []*Tree
			if found.Children() == nil {
				children = []*Tree{parent}
			} else if len(found.Children()) == 1 {
				// fmt.Println("len", len(found.Children()), "| first", found.Children()[0])
				last := false
				first := found.Children()[0]
				for !last {
					if first.Children() != nil {
						first = first.Children()[0]
					} else {
						first.children = []*Tree{parent}
						last = true
					}
				}

				children = found.Children()
			} else {
				// more than 1 child
				// fmt.Println("more than one child")
				children = append(found.Children(), parent)
			}
			// fmt.Println("---")
			root = New(found.Value())
			root.children = children
		}
	}

	return root
}

// PathTo returns the shortest path between two nodes in the tree.
func (tr *Tree) PathTo(from, to string) []string {
	panic("Please implement this function")
}
