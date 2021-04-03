class Stack {
  constructor() {
    this.data = [];
  }

  push(n) {
    this.data.push(n);
  }

  pop() {
    return this.data.pop();
  }

  cap() {
    return this.data.length;
  }

  empty() {
    return this.data.length === 0;
  }
  
}

module.exports = Stack;