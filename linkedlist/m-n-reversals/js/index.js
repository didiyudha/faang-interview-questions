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

  if (m >= n) {
    return
  }

  let lowerBoundIndex = m;
  let node = linkedList.getAt(m);
  let prev = linkedList.getAt(m-1);

  while(m <= n) {
    const nextNode = node.next;
    node.next = prev;
    prev = node;
    node = nextNode;
    m++;
  }
  
  linkedList.getAt(lowerBoundIndex-1).next = prev;
  linkedList.getAt(n).next = node;
}

module.exports = {
  Node,
  LinkedList,
  MNReversals,
}