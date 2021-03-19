package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList(head int) *LinkedList {
	node := Node{head, nil}
	linkedList := LinkedList{&node}
	return &linkedList
}

func (l *LinkedList) Add(data int) {
	node := l.Head

	if node == nil {
		l.Head = &Node{Data: data}
		return
	}

	if node.Next == nil {
		l.Head.Next = &Node{Data: data}
		return
	}

	for node.Next != nil {
		node = node.Next
	}

	node.Next = &Node{Data: data}

}

func (l *LinkedList) GetAt(index int) *Node {
	if index == 0 {
		if l.Head == nil {
			return nil
		}
		return l.Head
	}

	var counter int
	node := l.Head

	for node != nil {
		if counter == index {
			return node
		}
		counter++
		node = node.Next
	}

	return nil

}

func (l *LinkedList) Show() {
	node := l.Head
	for node != nil {
		fmt.Printf("%d, ", node.Data)
		node = node.Next
	}
	fmt.Println("")
}

func TraverseMN(linkedList *LinkedList, m, n int) {
	currentNode := linkedList.Head
	tail := linkedList.Head
	currentPointer := 1

	var prevBeforeStart *Node

	for currentPointer < m {
		prevBeforeStart = currentNode
		tail = currentNode.Next
		currentNode = currentNode.Next
		currentPointer++
	}

	var tempNode *Node

	for currentPointer >= m && currentPointer <= n {
		next := currentNode.Next
		currentNode.Next = tempNode
		tempNode = currentNode
		currentNode = next
		currentPointer++
	}

	prevBeforeStart.Next = tempNode
	tail.Next = currentNode

}

func main() {

}
