const Queue = require('./index');

test('test queue', () => {
  const queue = new Queue();

  queue.enqueue(1);
  queue.enqueue(2);
  queue.enqueue(3);
  queue.enqueue(4);
  queue.enqueue(5);

  expect(queue.peek()).toEqual(1);

  expect(queue.dequeue()).toEqual(1);
  expect(queue.dequeue()).toEqual(2);
  expect(queue.dequeue()).toEqual(3);
  expect(queue.dequeue()).toEqual(4);
  expect(queue.dequeue()).toEqual(5);
});