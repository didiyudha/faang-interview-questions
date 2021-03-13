const { isPalindrome } = require('./index');

test("should return true", () => {
  const s = "A man, a plan, a canal: Panama";
  const valid = isPalindrome(s);
  expect(valid).toEqual(true);
});