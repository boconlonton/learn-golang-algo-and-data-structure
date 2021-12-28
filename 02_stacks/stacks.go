package _2_stacks

type Node struct {
	Info string
	Next *Node
}

// Stack struct represent a stack data structure.
type Stack struct {
	First *Node
}

// IsEmpty method checks whether the stack is empty or not.
func (s *Stack) IsEmpty() bool {
	return s.First == nil
}

// Push method insert a new string onto stack
func (s *Stack) Push(data string) {
	f := s.First
	s.First = &Node{Info: data, Next: f}
}

// Pop method remove and return the string most recently added
func (s *Stack) Pop() string {
	f := s.First
	s.First = f.Next
	return f.Info
}
