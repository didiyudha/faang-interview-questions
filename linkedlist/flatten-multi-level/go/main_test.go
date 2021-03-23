package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestFlatten(t *testing.T) {

	var one, two, three, four, five, six, seven, eight, nine, ten, eleven, twelve, thirteen Node

	one.Value = 1
	one.Next = &two

	two.Value = 2
	two.Prev = &one
	two.Next = &three
	two.Child = &seven

	three.Value = 3
	three.Prev = &two
	three.Next = &four

	four.Value = 4
	four.Prev = &three
	four.Next = &five

	five.Value = 5
	five.Prev = &four
	five.Next = &six
	five.Child = &twelve

	six.Value = 6
	six.Prev = &five

	seven.Value = 7
	seven.Next = &eight

	eight.Value = 8
	eight.Next = &nine
	eight.Prev = &seven
	eight.Child = &ten

	nine.Value = 9
	nine.Prev = &eight

	twelve.Value = 12
	twelve.Next = &thirteen
	thirteen.Prev = &twelve

	ten.Value = 10
	ten.Next = &eleven

	eleven.Value = 11
	eleven.Prev = &ten

	twelve.Value = 12
	twelve.Next = &thirteen

	thirteen.Value = 13
	thirteen.Prev = &twelve

	Flatten(&one)

	expectedValueResult := []int{1, 2, 7, 8, 10, 11, 9, 3, 4, 5, 12, 13, 6}

	node := &one

	for i := 0; i < len(expectedValueResult); i++ {
		want := expectedValueResult[i]
		got := node.Value
		assert.Equal(t, want, got)
		node = node.Next
	}

}
