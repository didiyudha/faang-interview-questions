package main

import "sort"

func KthLargest(arr []int, k int) int {
	sort.Ints(arr)
	idxToFind := len(arr) - k
	return arr[idxToFind]
}
