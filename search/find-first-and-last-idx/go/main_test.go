package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchRange(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8

	result := searchRange(nums, target)
	assert.Equal(t, 3, result[0])
	assert.Equal(t, 4, result[1])

	target = 6
	result = searchRange(nums, target)
	assert.Equal(t, -1, result[0])
	assert.Equal(t, -1, result[1])
}
