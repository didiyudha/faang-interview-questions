package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidParenthesis(t *testing.T) {
	s := "({[]})"
	result := IsValidParenthesis(s)
	assert.True(t, result)

	s = "({[]])"
	result = IsValidParenthesis(s)
	assert.False(t, result)
}