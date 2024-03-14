package Stack

import "testing"

func TestPush(t *testing.T) {
	stack := Stack[int]{}
	stack.Push(1)

	if stack.Length() != 1 {
		t.Errorf("Expected length to be 1, got %v", stack.Length())
	}
	if stack.Items[0] != 1 {
		t.Errorf("Expected first item to be 1, got %v", stack.Items[0])
	}

	stack.Push(2)
	if stack.Length() != 2 {
		t.Errorf("Expected length to be 2, got %v", stack.Length())
	}
	if stack.Items[1] != 2 {
		t.Errorf("Expected second item to be 2, got %v", stack.Items[1])
	}

}

func TestPop(t *testing.T) {
	stack := Stack[int]{}
	stack.Push(1)
	stack.Push(2)

	item, err := stack.Pop()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if item != 2 {
		t.Errorf("Expected popped item to be 2, got %v", item)
	}
	if stack.Length() != 1 {
		t.Errorf("Expected length to be 1, got %v", stack.Length())
	}

	item, err = stack.Pop()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if item != 1 {
		t.Errorf("Expected popped item to be 1, got %v", item)
	}
	if stack.Length() != 0 {
		t.Errorf("Expected length to be 0, got %v", stack.Length())
	}

	_, err = stack.Pop()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestPeek(t *testing.T) {
	stack := Stack[int]{}
	_, err := stack.Peek()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	stack.Push(1)
	stack.Push(2)

	item, err := stack.Peek()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if item != 2 {
		t.Errorf("Expected peeked item to be 2, got %v", item)
	}
	if stack.Length() != 2 {
		t.Errorf("Expected length to be 2, got %v", stack.Length())
	}

	stack.Pop()
	item, err = stack.Peek()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if item != 1 {
		t.Errorf("Expected peeked item to be 1, got %v", item)
	}
	if stack.Length() != 1 {
		t.Errorf("Expected length to be 1, got %v", stack.Length())
	}

}

func TestContains(t *testing.T) {
	stack := Stack[int]{}
	stack.Push(1)
	stack.Push(2)

	if !stack.Contains(1) {
		t.Errorf("Expected to find value 1")
	}
	if stack.Contains(3) {
		t.Errorf("Expected to not find value 3")
	}
}
