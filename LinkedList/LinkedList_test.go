package LinkedList

import (
	"testing"
)

func TestAddNode(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	if list.Head.Value != 1 {
		t.Errorf("Expected head value to be 1, got %v", list.Head.Value)
	}
	list.Add(2)
	if list.Head.Next.Value != 2 {
		t.Errorf("Expected second node value to be 2, got %v", list.Head.Next.Value)
	}
	list.Add(3)
	if list.Length() != 3 {
		t.Errorf("Expected length to be 3, got %v", list.Length())
	}
}

func TestFind(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	node := list.find(1)
	if node == nil || node.Value != 1 {
		t.Errorf("Expected to find node with value 1, got %v", node)
	}
	node = list.find(3)
	if node != nil {
		t.Errorf("Expected to not find node, got %v", node)
	}
}

func TestIndexOf(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	if list.IndexOf(1) != 0 {
		t.Errorf("Expected index of 1 to be 0")
	}
	if list.IndexOf(2) != 1 {
		t.Errorf("Expected index of 2 to be 1")
	}
	if list.IndexOf(3) != -1 {
		t.Errorf("Expected index of 3 to be -1")
	}
}

func TestContains(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	if !list.Contains(1) {
		t.Errorf("Expected to find value 1")
	}
	if list.Contains(3) {
		t.Errorf("Expected to not find value 3")
	}
}

func TestRemove(t *testing.T) {
	list := LinkedList[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Remove(2)
	if list.Length() != 2 {
		t.Errorf("Expected length to be 2, got %v", list.Length())
	}
	if list.Head.Next.Value != 3 {
		t.Errorf("Expected second node value to be 3, got %v", list.Head.Next.Value)
	}
	list.Remove(1)
	if list.Length() != 1 {
		t.Errorf("Expected length to be 1, got %v", list.Length())
	}
	if list.Head.Value != 3 {
		t.Errorf("Expected head value to be 3, got %v", list.Head.Value)
	}
	list.Remove(3)
	if list.Length() != 0 {
		t.Errorf("Expected length to be 0, got %v", list.Length())
	}
	if list.Head != nil {
		t.Errorf("Expected head to be nil, got %v", list.Head)
	}
}
