class Node {
  constructor(data, next = null) {
    this.data = data;
    this.next = next;
  }
}

class LinkedList {
  constructor(head=null) {
    this.head = head;
  }

  add(data) {
    const node = new Node(data);

    if (!this.head) {
      this.head = node;
      return;
    }

    let lastNode = this.getLast();
    lastNode.next = node;
  }

  getLast() {
    let node = this.head;

    while(node.next) {
      node = node.next;
    }

    return node;

  }

  getAt(index) {
    if (index === 0) {
      return this.head;
    }
    let node = this.head;
    let counter = 0;
    while(node) {
      if (index === counter) {
        return node;
      }
      node = node.next;
      counter++;
    }

    return null;
  }
}

function reverseLinkedList(linkedList) {
  let previous = null;
  let node = linkedList.head;

  while(node) {
    const next = node.next;
    node.next = previous;
    previous = node;
    node = next;
  }

  linkedList.head = previous;
  
}

module.exports = {
  Node, 
  LinkedList,
  reverseLinkedList,
};