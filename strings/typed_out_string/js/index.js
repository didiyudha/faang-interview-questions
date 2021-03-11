function compareBackspaceString(S, T) {
  return build(S) === build(T);
}

function build(str) {
  const output = [];
  for (let s of str) {
    if (s !== '#') {
      output.push(s);
    } else if (output.length) {
      output.pop();
    }
  }
  return output.join('');
}

module.exports = {
  compareBackspaceString,
};