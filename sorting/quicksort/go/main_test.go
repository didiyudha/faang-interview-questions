package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	arr := []int{5, 3, 1, 6, 4, 2}
	QuickSort(arr, 0, len(arr)-1)

	assert.Equal(t, 1, arr[0])
}
