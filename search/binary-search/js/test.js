const binarySearch = require('./index');

test('binary search', () => {
  
  const nums = [1, 2, 3, 4, 5, 6, 7, 8, 9];
  const target = 8;

  const found = binarySearch(nums, target);
  expect(target).toEqual(found);

});