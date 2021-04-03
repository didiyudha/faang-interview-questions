const getPartition = function(nums, left, right) {
  let i = left;

  for (let j = left; j <= right; j++) {
    if (nums[j] <= nums[right]) {
      [nums[j], nums[i]] = [nums[i], nums[j]];
      i++;
    }
  }

  return i - 1;
}

function quickSort(arr, left, right) {
  if (left < right) {
    const partitionIdx = getPartition(arr, left, right);
    quickSort(arr, left, partitionIdx-1);
    quickSort(arr, partitionIdx+1, right);
  }
  
}

module.exports = quickSort;