const binarySearch = (nums, left, right, target) => {

  while (left <= right) {
    const middle = Math.floor((right+left) / 2);

    if (nums[middle] === target) {
      return middle;
    }

    if (nums[middle] < target) {
      left = middle+1;
    } else {
      right = middle - 1;
    }

  }

  return -1;

};

const searchRange = function (nums, target) {

  let left = 0;
  let right = nums.length-1;

  let firstPos = binarySearch(nums, left, right, target);
  if (firstPos === -1) {
    return [-1, -1];
  }

  let startPos = firstPos, endPos = firstPos;
  let startPosTemp, endPosTemp;

  while (startPos !== -1) {
    startPosTemp = startPos;
    startPos = binarySearch(nums, 0, startPos-1, target);
  }

  startPos = startPosTemp;

  while (endPos !== -1) {
    endPosTemp = endPos;
    endPos = binarySearch(nums, endPos+1, nums.length, target);
  }

  endPos = endPosTemp;

  return [startPos, endPos];

};

module.exports = searchRange;