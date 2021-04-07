package main

import "testing"

func TestQuickSelect(t *testing.T) {
	array := []int{5, 3, 1, 6, 4, 2}
	kToFind := 4
	n := kthLargest(array, kToFind)
	if n != 3 {
		t.Fatalf("Expected 3 but actual %v", n)
	}
}
