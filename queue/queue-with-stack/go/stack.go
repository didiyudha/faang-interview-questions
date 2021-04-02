package main

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{
		data: make([]int, 0),
	}
}

func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		return -1
	}
	el := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return el
}

func (s *Stack) Cap() int {
	return len(s.data)
}
