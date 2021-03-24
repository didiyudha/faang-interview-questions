package main

type Stack struct {
	elements []string
}

func NewStack() *Stack {
	elements := make([]string, 0)
	stack := Stack{elements: elements}
	return &stack
}

func (s *Stack) Push(val string) {
	s.elements = append(s.elements, val)
}

func (s *Stack) Pop() (val string) {

	if len(s.elements) == 0 {
		return ""

	}

	val = s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return val
}

func (s *Stack) Cap() int {
	return len(s.elements)
}