const Stack = require('./index');

describe('Stack', () => {
  test('should valid stack', () => {
    const s = new Stack();

    expect(s.empty()).toEqual(true);

    s.push(1);
    s.push(2);
    s.push(3);
    s.push(4);
    s.push(5);

    expect(s.cap()).toEqual(5);

    expect(s.pop()).toEqual(5);
    expect(s.pop()).toEqual(4);
    expect(s.pop()).toEqual(3);
    expect(s.pop()).toEqual(2);
    expect(s.pop()).toEqual(1);
  });
});