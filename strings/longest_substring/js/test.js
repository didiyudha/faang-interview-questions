const { longestSubstring, longestSubstrOptimal } = require('./index');

describe('Bruteforce soulution', () => {
  test('should return 4', () => {
    const s = "abcbdca";
    const total = longestSubstring(s);
    expect(total).toEqual(4);
  });
});

describe('Optimal soulution', () => {
  test('should return 4', () => {
    const s = "abcbdca";
    const total = longestSubstrOptimal(s);
    expect(total).toEqual(4);
  });
});