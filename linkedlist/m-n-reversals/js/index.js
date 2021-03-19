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

function MNReversals(linkedList, m, n) {

  let prevBeforeStart = null;
  let currentNode = linkedList.head;
  let tail = linkedList.head;
  let counter = 1;

  while(counter < m) {
    prevBeforeStart = currentNode;
    tail = tail.next;
    currentNode = currentNode.next;
    counter++;
  }

  let tempNode = null;

  while (counter >= m && counter <= n) {
    const next = currentNode.next;
    currentNode.next = tempNode;
    tempNode = currentNode;
    currentNode = next;
    counter++;
  }

  tail.next = currentNode;

  if (prevBeforeStart === null) {
    linkedList.head = tempNode
  } else {
    prevBeforeStart.next = tempNode;
  }
  

}

module.exports = {
  Node,
  LinkedList,
  MNReversals,
}