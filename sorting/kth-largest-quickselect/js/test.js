const findKthLargest = require('./index');

test('should return valid number', () => {
  const array = [5,3,1,6,4,2];
  const kToFind = 4;
  const n = findKthLargest(array, kToFind);
  expect(n).toEqual(3);
});