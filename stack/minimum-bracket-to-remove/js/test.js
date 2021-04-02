const minimumBracketToRemove = require('./index');

test('should return valid', () => {
  const s = "a)bc(de)"
  const result = minimumBracketToRemove(s);
  expect(result).toEqual('abc(de)');
});