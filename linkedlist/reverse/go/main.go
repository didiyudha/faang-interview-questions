package main

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Add(data interface{}) {

	if l.head == nil {
		l.head = &Node{
			data: data,
			next: nil,
		}
		return
	}

	node := l.head

	for node.next != nil {
		node = node.next
	}

	node.next = &Node{
		data: data,
		next: nil,
	}

}

func (l *LinkedList) GetAt(index int) *Node {

	if l.head == nil {
		return nil
	}

	node := l.head
	var counter int

	for node != nil {
		if index == counter {
			return node
		}
		node = node.next
		counter++
	}

	return nil
}

func ReverseLinkedList(linkedList *LinkedList) {
	if linkedList == nil || linkedList.head == nil || linkedList.head.next == nil {
		return
	}

	node := linkedList.head
	var prev *Node

	for node != nil {
		nextNode := node.next
		node.next = prev
		prev = node
		node = nextNode
	}

	linkedList.head = prev
}

func main() {

}
