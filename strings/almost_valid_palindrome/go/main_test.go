package main

import "testing"

func TestIsAlmostPalindrome(t *testing.T) {
	s := "abca"
	valid := IsAlmostPalindrome(s)
	if !valid {
		t.Fatalf("[FAIL] %s should almost palindrome", s)
	}
	t.Log("[PASS] Success")
}
