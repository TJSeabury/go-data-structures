package LinkedList

// Linked Lists: Linear collection of nodes where each node contains a value and a reference to the next node in the sequence.

type LLNode[T comparable] struct {
	Value T
	Next  *LLNode[T]
}

type LinkedList[T comparable] struct {
	Head *LLNode[T]
}

func New[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Length() int {
	count := 0
	currentNode := l.Head
	for currentNode != nil {
		count++
		currentNode = currentNode.Next
	}
	return count
}

func (l *LinkedList[T]) Add(value T) {
	newNode := &LLNode[T]{
		Value: value,
	}
	if l.Head == nil {
		l.Head = newNode
	} else {
		currentNode := l.Head
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		currentNode.Next = newNode
	}
}

func (l *LinkedList[T]) find(value T) *LLNode[T] {
	if l.Head == nil {
		return nil
	}
	currentNode := l.Head
	for currentNode != nil {
		if currentNode.Value == value {
			return currentNode
		} else {
			currentNode = currentNode.Next
		}
	}
	return nil
}

func (l *LinkedList[T]) Contains(value T) bool {
	return l.find(value) != nil
}

func (l *LinkedList[T]) IndexOf(value T) int {
	index := 0
	currentNode := l.Head
	for currentNode != nil {
		if currentNode.Value == value {
			return index
		} else {
			currentNode = currentNode.Next
			index++
		}
	}
	return -1
}

func (l *LinkedList[T]) InsertBefore(value T, index int) {
	if index == 0 {
		newNode := &LLNode[T]{
			Value: value,
			Next:  l.Head,
		}
		l.Head = newNode
		return
	}
	currentNode := l.Head
	for i := 0; i < index-1; i++ {
		if currentNode == nil {
			return
		}
		currentNode = currentNode.Next
	}
	newNode := &LLNode[T]{
		Value: value,
		Next:  currentNode.Next,
	}
	currentNode.Next = newNode
}

func (l *LinkedList[T]) InsertAfter(value T, index int) {
	currentNode := l.Head
	for i := 0; i < index; i++ {
		if currentNode == nil {
			return
		}
		currentNode = currentNode.Next
	}
	newNode := &LLNode[T]{
		Value: value,
		Next:  currentNode.Next,
	}
	currentNode.Next = newNode
}

func (l *LinkedList[T]) Remove(value T) {
	if l.Head == nil {
		return
	}
	if l.Head.Value == value {
		l.Head = l.Head.Next
		return
	}
	currentNode := l.Head
	for currentNode.Next != nil {
		if currentNode.Next.Value == value {
			currentNode.Next = currentNode.Next.Next
			return
		} else {
			currentNode = currentNode.Next
		}
	}
}
