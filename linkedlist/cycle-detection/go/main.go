package main

type Node struct {
	Value int
	Next *Node
}

func NewNode(val int, next *Node) *Node {
	return &Node{
		Value: val,
		Next:  next,
	}
}

func FindCyclicNode(head *Node) *Node {

	hare := head
	tortoise := head

	for {

		if hare.Next == nil || tortoise.Next == nil || tortoise.Next.Next == nil {
			return nil
		}

		hare = hare.Next
		tortoise = tortoise.Next.Next

		if hare == tortoise {
			break
		}

	}

	p1 := head
	p2 := tortoise

	for {

		p1 = p1.Next
		p2 = p2.Next

		if p1 == p2 {
			return p1
		}
	}

}

func main() {

}