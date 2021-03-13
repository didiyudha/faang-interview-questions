const { almostPalindrome } = require('./index');

test("should return true", () => {
  const s = "abca";
  const valid = almostPalindrome(s);
  expect(valid).toEqual(true);
});

test("should return true", () => {
  const s = "aba";
  const valid = almostPalindrome(s);
  expect(valid).toEqual(true);
});