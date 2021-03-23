package main

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
	Child *Node
}

func Flatten(head *Node) {

	node := head

	for node != nil {

		if node.Child == nil {
			node = node.Next
			continue
		}

		next := node.Next
		headOfChild := node.Child
		tail := node.Child

		for tail.Next != nil {
			tail = tail.Next
		}

		node.Next = headOfChild
		tail.Next = next
		node = node.Next

	}
}

func main() {

}
