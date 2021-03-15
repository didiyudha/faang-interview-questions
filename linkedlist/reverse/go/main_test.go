package main

import "testing"

func TestReverseLinkedList(t *testing.T) {

	linkedList := NewLinkedList()
	linkedList.Add(1)
	linkedList.Add(2)
	linkedList.Add(3)
	linkedList.Add(4)
	linkedList.Add(5)

	if linkedList.GetAt(0).data != 1 {
		t.Fatalf("Expecting 1 but actual %v", linkedList.GetAt(0).data)
	}
	if linkedList.GetAt(1).data != 2 {
		t.Fatalf("Expecting 2 but actual %v", linkedList.GetAt(0).data)
	}
	if linkedList.GetAt(2).data != 3 {
		t.Fatalf("Expecting 3 but actual %v", linkedList.GetAt(0).data)
	}
	if linkedList.GetAt(3).data != 4 {
		t.Fatalf("Expecting 4 but actual %v", linkedList.GetAt(0).data)
	}
	if linkedList.GetAt(4).data != 5 {
		t.Fatalf("Expecting 5 but actual %v", linkedList.GetAt(0).data)
	}

	ReverseLinkedList(linkedList)

	if linkedList.GetAt(0).data != 5 {
		t.Fatalf("Expecting 5 but actual %v", linkedList.GetAt(0).data)
	}
	if linkedList.GetAt(1).data != 4 {
		t.Fatalf("Expecting 4 but actual %v", linkedList.GetAt(0).data)
	}
	if linkedList.GetAt(2).data != 3 {
		t.Fatalf("Expecting 3 but actual %v", linkedList.GetAt(0).data)
	}
	if linkedList.GetAt(3).data != 2 {
		t.Fatalf("Expecting 2 but actual %v", linkedList.GetAt(0).data)
	}
	if linkedList.GetAt(4).data != 1 {
		t.Fatalf("Expecting 1 but actual %v", linkedList.GetAt(0).data)
	}

}
