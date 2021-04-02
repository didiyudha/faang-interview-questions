class Queue {
  constructor() {
    this.in = [];
    this.out = [];
  }

  enqueue(x) {
    this.in.push(x);
  }

  dequeue() {
    if (!this.out.length) {
      while(this.in.length) {
        this.out.push(this.in.pop());
      }
    }
    return this.out.pop();
  }

  peek() {
    if (!this.out.length) {
      while(this.in.length) {
        this.out.push(this.in.pop());
      }
    }
    return this.out[this.out.length-1];
  }

  isEmpty() {
    return this.in.length === 0 && this.out.length === 0    
  }
}

module.exports = Queue;