package Stack

import "errors"

// Stacks: Last-In-First-Out (LIFO) data structure.
var ErrEmptyStack = errors.New("empty stack")

type Stack[T comparable] struct {
	Items []T
}

func New[T comparable]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Length() int {
	return len(s.Items)
}

func (s *Stack[T]) Push(value T) {
	s.Items = append(s.Items, value)
}

func (s *Stack[T]) Pop() (T, error) {
	var zeroValue T
	if s.Length() == 0 {
		return zeroValue, ErrEmptyStack
	}
	lastIndex := s.Length() - 1
	lastItem := s.Items[lastIndex]
	s.Items = s.Items[:lastIndex]
	return lastItem, nil
}

func (s *Stack[T]) Peek() (T, error) {
	var zeroValue T
	if s.Length() == 0 {
		return zeroValue, ErrEmptyStack
	}
	return s.Items[s.Length()-1], nil
}

func (s *Stack[T]) Contains(value T) bool {
	for _, item := range s.Items {
		if item == value {
			return true
		}
	}
	return false
}
