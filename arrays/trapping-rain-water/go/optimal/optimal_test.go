package main

import "testing"

func TestTrapWater(t *testing.T) {

	tableTest := []struct {
		Input  []int
		Result int
	}{
		{
			Input:  []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			Result: 6,
		},
		{
			Input:  []int{4, 2, 0, 3, 2, 5},
			Result: 9,
		},
	}

	for _, tt := range tableTest {
		total := TrapWater(tt.Input)
		if total != tt.Result {
			t.Fatalf("[FAIL] actual %d expected %d\n", total, tt.Result)
		}
		t.Log("[PASS] Success")
	}
}
