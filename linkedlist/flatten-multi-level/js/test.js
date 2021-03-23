const { TestScheduler } = require('@jest/core');
const { ListNode, flatten } = require('./index');

const nodes = [1, [2, 7, [8, 10, 11], 9], 3, 4, [5, 12, 13], 6]

const buildMultiLevel = function(nodes) {
  const head = new ListNode(nodes[0]);
  let currentNode = head;

  for(let i = 1; i < nodes.length; i++) {
    const val = nodes[i];
    let nextNode;

    if(Array.isArray(val)) {
      nextNode = buildMultiLevel(val);
      currentNode.child = nextNode;
      continue;
    }

    nextNode = new ListNode(val);
    currentNode.next = nextNode;
    nextNode.prev = currentNode;
    currentNode = nextNode;
  }
  
  return head;
}



test('flatten multi level doubly linked list', () => {
  let multiLinkedList = buildMultiLevel(nodes);
  let node = flatten(multiLinkedList);

  const resultVal = [1, 2, 7, 8, 10, 11, 9, 3, 4, 5, 12, 13, 6];

  for (let i = 0; i < resultVal.length; i++) {
    expect(node.val).toEqual(resultVal[i]);
    node = node.next;
  }

});