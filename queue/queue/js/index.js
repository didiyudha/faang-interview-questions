class Queue {
  constructor() {
    this.data = [];
  }

  enqueue(x) {
    this.data.unshift(x);
  }

  dequeue() {
    return this.data.pop();
  }

  cap() {
    return this.data.length;
  }

  empty() {
    return this.data.length === 0;
  }

}

module.exports = Queue;