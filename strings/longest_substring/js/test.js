const { longestSubstring } = require('./index');

describe('Bruteforce soulution', () => {
  test('should return 3', () => {
    const s = "abcbdca";
    const total = longestSubstring(s);
    expect(total).toEqual(4);
  });
});