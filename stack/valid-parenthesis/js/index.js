const storage = {
  '(': ')',
  '{': '}',
  '[': ']',
};

function isValidParenthesis(s) {

  const stack = [];

  for (let i = 0; i < s.length; i++) {
    
    const char = s[i];
    const closeParenExist = storage[char];

    if (closeParenExist) {
      stack.push(char);
      continue;
    }

    if (stack.length === 0) {
      return false;
    }

    const lastChar = stack.pop();
    const correctParen = storage[lastChar];

    if (char !== correctParen) {
      return false;
    }

  }

  return stack.length === 0;

}

module.exports = isValidParenthesis;