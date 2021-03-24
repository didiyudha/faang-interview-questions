const { ListNode, findCyclePoint } = require('./index');

describe('when we want to check if linked list is cyclic or not', () => {
  test('findCyclePoint function should be able to detect it ', () => {
    const linkedList = [8,7,6,5,4,3,2,1].reduce((acc, val) => new ListNode(val, acc), null);

      let curr = linkedList, cycleNode;
      while(curr.next !== null) {
        if(curr.val === 3) {
          cycleNode = curr;
        }

        curr = curr.next;
      }

      curr.next = cycleNode;

      const node = findCyclePoint(linkedList);
      
      expect(node.val).toEqual(3)
  })
});