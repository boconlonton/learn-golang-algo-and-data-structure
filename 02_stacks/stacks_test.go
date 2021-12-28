package _2_stacks

import "testing"

func TestStacks(t *testing.T) {
	s := Stack{First: nil}
	checkEmpty := s.IsEmpty()
	if !checkEmpty {
		t.Errorf("expected true but got %v.", checkEmpty)
	}
	s.Push("World!")
	s.Push("Hello")
	r := s.Pop()
	if r != "Hello" {
		t.Errorf("expected `Hello` but got %v.", r)
	}
}
