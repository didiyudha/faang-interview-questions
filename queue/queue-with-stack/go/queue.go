package main

type Queue struct {
	In  *Stack
	Out *Stack
}

func NewQueue() *Queue {
	in := NewStack()
	out := NewStack()
	return &Queue{
		In:  in,
		Out: out,
	}
}

func (q *Queue) Enqueue(x int) {
	q.In.Push(x)
}

func (q *Queue) Dequeue() int {
	if q.Out.Cap() == 0 {
		for q.In.Cap() > 0 {
			q.Out.Push(q.In.Pop())
		}
	}

	return q.Out.Pop()
}

func (q *Queue) Peek() int {
	if q.Out.Cap() == 0 {
		for q.In.Cap() > 0 {
			q.Out.Push(q.In.Pop())
		}
	}
	return q.Out.data[len(q.Out.data)-1]
}

func (q *Queue) Empty() bool {
	return q.In.Cap() == 0 && q.Out.Cap() == 0
}
