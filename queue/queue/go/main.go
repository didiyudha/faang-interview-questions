package main

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return &Queue{
		data: make([]int, 0),
	}
}

func (q *Queue) Enqueue(x int) {
	q.data = append([]int{x}, q.data...)
}

func (q *Queue) Dequeue() int {
	dq := q.data[q.Size()-1]
	q.data = q.data[:q.Size()-1]
	return dq
}

func (q *Queue) Size() int {
	return len(q.data)
}

func main() {

}
