const { compareBackspaceString } = require('./index');

describe('Bruteforce', () => {

  test('should return true with string ac', () => {
    const S = "ab#c";
    const T = "ad#c";
    const isEqual = compareBackspaceString(S, T);
    expect(isEqual).toEqual(true);
  })

  test('should return true with empty string', () => {
    const S = "ab##";
    const T = "c#d#";
    const isEqual = compareBackspaceString(S, T);
    expect(isEqual).toEqual(true);
  });

  test('should return true with string c', () => {
    const S = "a##c";
    const T = "#a#c";
    const isEqual = compareBackspaceString(S, T);
    expect(isEqual).toEqual(true);
  });

  test('should return false', () => {
    const S = "a#c";
    const T = "b";
    const isEqual = compareBackspaceString(S, T);
    expect(isEqual).toEqual(false);
  });

});