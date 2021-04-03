package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue()
	assert.NotNil(t, q)
	assert.Equal(t, 0, q.Size())
}

func TestEnqueue(t *testing.T) {
	q := NewQueue()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	assert.Equal(t, 5, q.Size())
}

func TestDequeue(t *testing.T) {
	q := NewQueue()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)

	var dq int

	for j := 1; j <= 5; j++ {
		dq = q.Dequeue()
		assert.Equal(t, j, dq)
	}

}
