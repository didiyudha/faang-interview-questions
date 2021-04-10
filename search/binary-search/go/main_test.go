package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {

	nums := []int{1, 2, 3, 4, 5 , 6, 7, 8, 9}

	for _, v := range nums {
		foundNumber := binarySearch(nums, v)
		assert.Equal(t, v,  foundNumber)
	}

	target := 100
	foundNumber := binarySearch(nums, target)
	assert.NotEqual(t, target, foundNumber)

	target = 101
	foundNumber = binarySearch(nums, target)
	assert.NotEqual(t, target, foundNumber)
}