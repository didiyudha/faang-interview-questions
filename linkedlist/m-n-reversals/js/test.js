const { Node, LinkedList, MNReversals } = require('./index');

describe('MN linked list reversals', () => {
  test('should reversals linked list', () => {
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

    MNReversals(linkedList, 2, 4);

    expect(linkedList.getAt(0).data).toEqual(1);
    expect(linkedList.getAt(1).data).toEqual(4);
    expect(linkedList.getAt(2).data).toEqual(3);
    expect(linkedList.getAt(3).data).toEqual(2);
    expect(linkedList.getAt(4).data).toEqual(5);
  });

  test('should reversals linked list', () => {
    const head = new Node(1);
    const linkedList = new LinkedList(head);
    linkedList.add(2);
    linkedList.add(3);
    linkedList.add(4);
    linkedList.add(5);
    linkedList.add(6);
    linkedList.add(7);
    linkedList.add(8);
    linkedList.add(9);
    linkedList.add(10);

    MNReversals(linkedList, 5, 7);

    expect(linkedList.getAt(0).data).toEqual(1);
    expect(linkedList.getAt(1).data).toEqual(2);
    expect(linkedList.getAt(2).data).toEqual(3);
    expect(linkedList.getAt(3).data).toEqual(4);
    expect(linkedList.getAt(4).data).toEqual(7);
    expect(linkedList.getAt(5).data).toEqual(6);
    expect(linkedList.getAt(6).data).toEqual(5);
    expect(linkedList.getAt(7).data).toEqual(8);
    expect(linkedList.getAt(8).data).toEqual(9);
    expect(linkedList.getAt(9).data).toEqual(10);


  });

  test('should reversals linked list', () => {
    const head = new Node(1);
    const linkedList = new LinkedList(head);
    linkedList.add(2);
    linkedList.add(3);
    linkedList.add(4);
    linkedList.add(5);
    linkedList.add(6);
    linkedList.add(7);
    linkedList.add(8);
    linkedList.add(9);
    linkedList.add(10);

    MNReversals(linkedList, 1, 2);

    expect(linkedList.getAt(0).data).toEqual(2);
    expect(linkedList.getAt(1).data).toEqual(1);
    expect(linkedList.getAt(2).data).toEqual(3);
    expect(linkedList.getAt(3).data).toEqual(4);
    expect(linkedList.getAt(4).data).toEqual(5);
    expect(linkedList.getAt(5).data).toEqual(6);
    expect(linkedList.getAt(6).data).toEqual(7);
    expect(linkedList.getAt(7).data).toEqual(8);
    expect(linkedList.getAt(8).data).toEqual(9);
    expect(linkedList.getAt(9).data).toEqual(10);
  });

  test('should reversals linked list', () => {
    const head = new Node(1);
    const linkedList = new LinkedList(head);
    linkedList.add(2);
    linkedList.add(3);
    linkedList.add(4);
    linkedList.add(5);
    linkedList.add(6);
    linkedList.add(7);
    linkedList.add(8);
    linkedList.add(9);
    linkedList.add(10);

    MNReversals(linkedList, 1, 10);

    expect(linkedList.getAt(0).data).toEqual(10);
    expect(linkedList.getAt(1).data).toEqual(9);
    expect(linkedList.getAt(2).data).toEqual(8);
    expect(linkedList.getAt(3).data).toEqual(7);
    expect(linkedList.getAt(4).data).toEqual(6);
    expect(linkedList.getAt(5).data).toEqual(5);
    expect(linkedList.getAt(6).data).toEqual(4);
    expect(linkedList.getAt(7).data).toEqual(3);
    expect(linkedList.getAt(8).data).toEqual(2);
    expect(linkedList.getAt(9).data).toEqual(1);
  });
})