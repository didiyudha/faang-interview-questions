function longestSubstring(s) {

  let longest = 0;

  for (let i = 0; i < s.length; i++) {

    let m = {};
    let currentLength = 0;

    for (let j = i; j < s.length; j++) {
      const exists = m[s[j]];

      if (exists) {
        break;
      }

      currentLength++;
      m[s[j]] = true;
      longest = Math.max(longest, currentLength);

    }
  }

  return longest
}

module.exports = {
  longestSubstring,
};