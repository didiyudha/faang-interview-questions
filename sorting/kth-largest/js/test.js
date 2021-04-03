const kthLargest = require('./index');

test('should return kth largest of the data', () => {
  const arr = [5, 3, 1, 6, 4, 2];
  const k = 4;

  const n = kthLargest(arr, k);
  expect(n).toEqual(3);
})