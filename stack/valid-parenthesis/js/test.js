const isValidParenthesis = require('./index');

describe('Given a string [{()}]', () => {
  test('It should return true', () => {
    const s = '[{()}]';
    const result = isValidParenthesis(s);
    expect(result).toEqual(true);
  });
})