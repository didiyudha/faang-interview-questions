package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	stack := NewStack()
	assert.NotNil(t, stack)
	assert.Equal(t, 0, len(stack.data))
}

func TestCap(t *testing.T) {
	stack := NewStack()

	cap := 5

	for i := 1; i <= cap; i++ {
		stack.Push(i)
	}

	assert.Equal(t, cap, stack.Cap())
}

func TestPush(t *testing.T) {
	stack := NewStack()

	cap := 5

	for i := 1; i <= cap; i++ {
		stack.Push(i)
	}

	assert.Equal(t, cap, stack.Cap())
}

func TestPop(t *testing.T) {
	stack := NewStack()

	cap := 5

	for i := 1; i <= cap; i++ {
		stack.Push(i)
	}

	for j := cap; j >= 1; j-- {
		number := stack.Pop()
		assert.Equal(t, j, number)
	}
}
