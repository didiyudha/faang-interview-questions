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

function longestSubstrOptimal(s) {

  let left = 0, total = 0;
  let m = new Map();

  for (let right = 0; right < s.length; right++) {
      const character = s[right];
      const prevSeen = m.get(character);

      if (prevSeen >= left) {
          left = prevSeen + 1;
      }

      m.set(character, right);
      const currentLen = right - left + 1;
      total = Math.max(currentLen, total);
  }

  return total;
  
};


module.exports = {
  longestSubstring,
  longestSubstrOptimal,
};