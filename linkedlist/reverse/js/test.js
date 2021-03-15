const { Node, LinkedList, reverseLinkedList } = require('./index');

describe('Reverse Linked List', () => {
  test('should return reverse linked list', () => {
    const head = new Node(1);
    const linkedList = new LinkedList(head);
    linkedList.add(2);
    linkedList.add(3);
    linkedList.add(4);
    linkedList.add(5);

    expect(linkedList.getAt(0).data).toEqual(1);
    expect(linkedList.getAt(1).data).toEqual(2);
    expect(linkedList.getAt(2).data).toEqual(3);
    expect(linkedList.getAt(3).data).toEqual(4);
    expect(linkedList.getAt(4).data).toEqual(5);

    reverseLinkedList(linkedList);

    expect(linkedList.getAt(0).data).toEqual(5);
    expect(linkedList.getAt(1).data).toEqual(4);
    expect(linkedList.getAt(2).data).toEqual(3);
    expect(linkedList.getAt(3).data).toEqual(2);
    expect(linkedList.getAt(4).data).toEqual(1);

  });
})