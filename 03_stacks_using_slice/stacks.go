package _3_stacks_using_slice

type Stack []string

// IsEmpty method check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push method insert new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

// Pop method remove and return the last item
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		i := len(*s) - 1
		e := (*s)[i]
		*s = (*s)[:i]
		return e, true
	}
}
