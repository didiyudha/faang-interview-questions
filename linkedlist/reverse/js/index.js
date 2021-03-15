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
  let node = linkedList.head;
  let prev = null;

  while(node) {
    const nextNode = node.next;
    node.next = prev;
    prev = node;
    node = nextNode;
  }

  linkedList.head = prev;
}

module.exports = {
  Node, 
  LinkedList,
  reverseLinkedList,
};