function almostPalindrome(s) {
  let left = 0, right = s.length - 1;

  while(left <= right) {

    if (s[left] !== s[right]) {
      return almostValidSubPalindrome(s, left+1, right) || almostValidSubPalindrome(s, left, right-1);
    }

    left++;
    right--;

  }

  return true;
}

function almostValidSubPalindrome(s, left, right) {
  while(left <= right) {
    if (s[left] !== s[right]) {
      return false;
    }
    left++;
    right--;
  }
  return true;
}

module.exports = {
  almostPalindrome,
};