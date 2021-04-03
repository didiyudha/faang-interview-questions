const Queue = require('./index');

describe('Queue', () => {
  test('Should valid queue', () => {
    const q = new Queue();

    for (let i = 1; i <= 5; i++) {
      q.enqueue(i);
    }

    expect(q.cap()).toEqual(5);
    expect(q.empty()).toEqual(false);

    for (let i = 1; i <= 5; i++) {
      const dq = q.dequeue()
      expect(dq).toEqual(i);
    }

    expect(q.empty()).toEqual(true);

  })
})