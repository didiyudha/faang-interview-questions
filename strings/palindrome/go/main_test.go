package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	s := "saras"
	valid := IsPalindrome(s)
	if !valid {
		t.Fatalf("[FAIL] %s should palindrome\n", s)
	}
	t.Log("[PASS] Success")
}
