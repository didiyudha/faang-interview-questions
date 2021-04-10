function binarySearch(nums, target) {

  let left = 0, right = nums.length-1;
  let middle = (right+left) / 2;

  while (left <= right) {

    if (nums[middle] === target) {
      return nums[middle];
    }

    if (nums[middle] < target) {
      left = middle + 1;
      middle = Math.ceil((right + left) / 2);
      continue;
    }
    
    right = middle - 1;
    middle = Math.ceil((right + left) / 2);

  }

  return - 1;

}

module.exports = binarySearch;