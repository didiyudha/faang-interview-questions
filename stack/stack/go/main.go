package main

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{
		data: make([]int, 0),
	}
}

func (s *Stack) Push(n int) {
	s.data = append(s.data, n)
}

func (s *Stack) Cap() int {
	return len(s.data)
}

func (s *Stack) Pop() int {
	n := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return n
}

func main() {

}
