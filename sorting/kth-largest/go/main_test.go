package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthLargest(t *testing.T) {
	arr := []int{5, 3, 1, 6, 4, 2}
	kth := 4

	// 1, 2, 3, 4, 5, 6
	//			 4	3	 2	1

	n := KthLargest(arr, kth)
	assert.Equal(t, 3, n)
}
