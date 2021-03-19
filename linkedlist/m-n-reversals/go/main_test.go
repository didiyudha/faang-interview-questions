package main

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestTraverseMN(t *testing.T) {
	t.Run("First case", func(t *testing.T) {
		linkedList := NewLinkedList(1)
		linkedList.Add(2)
		linkedList.Add(3)
		linkedList.Add(4)
		linkedList.Add(5)

		TraverseMN(linkedList, 2, 4)

		assert.Equal(t, linkedList.GetAt(0).Data, 1)
		assert.Equal(t, linkedList.GetAt(1).Data, 4)
		assert.Equal(t, linkedList.GetAt(2).Data, 3)
		assert.Equal(t, linkedList.GetAt(3).Data, 2)
		assert.Equal(t, linkedList.GetAt(4).Data, 5)
	})

	t.Run("Second case", func(t *testing.T) {
		linkedList := NewLinkedList(1)
		linkedList.Add(2)
		linkedList.Add(3)
		linkedList.Add(4)
		linkedList.Add(5)
		linkedList.Add(6)
		linkedList.Add(7)
		linkedList.Add(8)
		linkedList.Add(9)
		linkedList.Add(10)

		TraverseMN(linkedList, 4, 8)

		assert.Equal(t, linkedList.GetAt(0).Data, 1)
		assert.Equal(t, linkedList.GetAt(1).Data, 2)
		assert.Equal(t, linkedList.GetAt(2).Data, 3)
		assert.Equal(t, linkedList.GetAt(3).Data, 8)
		assert.Equal(t, linkedList.GetAt(4).Data, 7)
		assert.Equal(t, linkedList.GetAt(5).Data, 6)
		assert.Equal(t, linkedList.GetAt(6).Data, 5)
		assert.Equal(t, linkedList.GetAt(7).Data, 4)
		assert.Equal(t, linkedList.GetAt(8).Data, 9)
		assert.Equal(t, linkedList.GetAt(9).Data, 10)

	})
}
