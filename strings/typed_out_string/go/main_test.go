package main

import (
	"testing"
)

func TestBackspaceCompare(t *testing.T) {
	tableTest := []struct {
		S      string
		T      string
		Result bool
	}{
		{
			S:      "ab#c",
			T:      "ad#c",
			Result: true,
		},
		{
			S:      "ab##",
			T:      "c#d#",
			Result: true,
		},
		{
			S:      "a##c",
			T:      "#a#c",
			Result: true,
		},
		{
			S:      "a#c",
			T:      "b",
			Result: false,
		},
	}

	for _, tt := range tableTest {
		result := backspaceCompare(tt.S, tt.T)
		if result != tt.Result {
			t.Fatalf("Expected %v actual %v", tt.Result, result)
		}
		t.Log("[PASS] Success")
	}
}

func TestBuild(t *testing.T) {

	tableTest := []struct {
		Str    string
		Result string
	}{
		{
			Str:    "ab#c",
			Result: "ac",
		},
		{
			Str:    "ad#c",
			Result: "ac",
		},
		{
			Str:    "ab##",
			Result: "",
		},
		{
			Str:    "c#d#",
			Result: "",
		},
		{
			Str:    "a##c",
			Result: "c",
		},
		{
			Str:    "#a#c",
			Result: "c",
		},
		{
			Str:    "a#c",
			Result: "c",
		},
	}

	for _, tt := range tableTest {
		result := build(tt.Str)
		if result != tt.Result {
			t.Fatalf("[FAIL] expected %v actual %v\n", tt.Result, result)
		}
		t.Logf("[PASS] Success %s => %s", tt.Str, tt.Result)
	}
}

func TestBackspaceCompareOptimal(t *testing.T) {
	tableTest := []struct {
		S      string
		T      string
		Result bool
	}{
		{
			S:      "ab#c",
			T:      "ad#c",
			Result: true,
		},
		{
			S:      "ab##",
			T:      "c#d#",
			Result: true,
		},
		{
			S:      "a##c",
			T:      "#a#c",
			Result: true,
		},
		{
			S:      "a#c",
			T:      "b",
			Result: false,
		},
	}

	for _, tt := range tableTest {
		result := backspaceCompareOptimal(tt.S, tt.T)
		if result != tt.Result {
			t.Fatalf("Expected %v actual %v", tt.Result, result)
		}
		t.Log("[PASS] Success")
	}
}
