class ListNode {
  constructor(val, next = null, prev = null, child = null) {
    this.val = val;
    this.next = next;
    this.prev = prev;
    this.child = child;
  }
}

function flatten(head) {
  let node = head;

  while(node) {

    if (!node.child) {
      node = node.next;
      continue;
    }

    const next = node.next;

    let childHead = node.child;
    let tail = node.child;

    while(tail.next) {
      tail = tail.next;
    }

    node.next = childHead;
    childHead.prev = node;
    node.child = null;

    if (next) {
      tail.next = next;
      next.prev = tail;
    }

  }

  return head;

}

module.exports = {
  ListNode,
  flatten,
};