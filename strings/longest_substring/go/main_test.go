package main

import "testing"

func TestLongestSubstr(t *testing.T) {
	s := "abcbdca"
	total := longestSubstr(s)
	if total != 4 {
		t.Fatalf("[FAIL] total should 3 but actual %d", total)
	}
	t.Log("[PASS] Success")
}

func TestLongestSubstrOptimal(t *testing.T) {
	s := "abcbdca"
	total := longestSubstrOptimal(s)
	if total != 4 {
		t.Fatalf("[FAIL] total should 3 but actual %d", total)
	}
	t.Log("[PASS] Success")
}
