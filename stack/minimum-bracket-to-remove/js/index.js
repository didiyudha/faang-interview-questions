function minimumBracketToRemove(s) {
  const arr = s.split('');
  const stack = [];

  for (let i = 0; i < s.length; i++) {
    const c = s[i];
    if (c === '(') {
      stack.push(i);
      continue;
    }
    if (c === ')' && stack.length) {
      stack.pop();
      continue;
    }
    if (c === ')') {
      arr[i] = '';
    }
  }

  while(stack.length) {
    const idx = stack.pop();
    arr[idx] = '';
  }

  return arr.join('');

}

module.exports = minimumBracketToRemove;