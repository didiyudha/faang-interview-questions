package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func generateCyclicLinkedList() *Node {

	arr := []int{8, 7, 6, 5, 4, 3, 2, 1}

	var cyclicNode *Node
	var next *Node

	head := NewNode(arr[0], next)
	node := head

	for i := 1; i < len(arr); i++ {
		next := NewNode(arr[i], nil)
		node.Next = next
		node = node.Next

		if arr[i] == 3 {
			cyclicNode = node
		}
	}

	node.Next = cyclicNode

	return head
}

func TestFindCyclicNode(t *testing.T) {
	head := generateCyclicLinkedList()

	node := FindCyclicNode(head)
	assert.Equal(t, 3, node.Value)

}
