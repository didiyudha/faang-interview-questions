package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	assert.Equal(t, 1, q.Peek())
	out := q.Dequeue()
	assert.Equal(t, 1, out)

	assert.Equal(t, 2, q.Dequeue())
	assert.Equal(t, 3, q.Dequeue())
	assert.Equal(t, 4, q.Dequeue())
	assert.Equal(t, 5, q.Dequeue())

	assert.True(t, q.Empty())
}
