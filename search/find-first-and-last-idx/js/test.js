const searchRange = require('./index');

test('find first and las position element in sorted array', () => {
  const nums = [5,7,7,8,8,10];
  const target = 8;

  const result = searchRange(nums, target);
  expect(result[0]).toEqual(3);
  expect(result[1]).toEqual(4);

})