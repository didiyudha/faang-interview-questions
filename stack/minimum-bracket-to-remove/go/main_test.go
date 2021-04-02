package main

import "testing"

func TestMinimumBracket(t *testing.T) {
	s := "))(("
	expectedResult := ""
	result := MinimumBracket(s)

	if result != expectedResult {
		t.Fatalf("Expected: %v Actual: %v\n", expectedResult, result)
	}

}
