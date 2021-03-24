package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack()
	assert.NotNil(t, stack)
	assert.Equal(t, 0, len(stack.elements))
}

func TestPush(t *testing.T) {
	stack := NewStack()
	stack.Push("1")
	assert.Equal(t, 1, len(stack.elements))
	assert.Equal(t, "1", stack.elements[0])
}

func TestPop(t *testing.T) {
	stack := NewStack()
	popVal := stack.Pop()
	assert.Equal(t, "", popVal)

	stack.Push("1")
	stack.Push("2")
	stack.Push("3")
	stack.Push("4")
	stack.Push("5")

	popVal = stack.Pop()
	assert.Equal(t, "5", popVal)

	popVal = stack.Pop()
	assert.Equal(t, "4", popVal)

	popVal = stack.Pop()
	assert.Equal(t, "3", popVal)

	popVal = stack.Pop()
	assert.Equal(t, "2", popVal)

	popVal = stack.Pop()
	assert.Equal(t, "1", popVal)


	popVal = stack.Pop()
	assert.Equal(t, "", popVal)
}

func TestCap(t *testing.T) {
	stack := NewStack()
	cap := stack.Cap()
	assert.Equal(t, 0, cap)

	stack.Push("1")
	stack.Push("2")
	stack.Push("3")
	stack.Push("4")
	stack.Push("5")

	cap = stack.Cap()
	assert.Equal(t, 5, cap)
}