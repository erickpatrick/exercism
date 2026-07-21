package listops

import "slices"

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, value := range s {
		initial = fn(initial, value)
	}

	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	s = s.Reverse()
	for _, value := range s {
		initial = fn(value, initial)
	}

	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	var result IntList

	for _, value := range s {
		if fn(value) {
			result = append(result, value)
		}
	}

	return result
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	var result IntList

	for _, value := range s {
		result = append(result, fn(value))
	}

	return result
}

func (s IntList) Reverse() IntList {
	list := s
	slices.Reverse(list)

	return list
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, list := range lists {
		s = s.Append(list)
	}

	return s
}
