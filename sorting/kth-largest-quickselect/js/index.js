function getPartitionIndex(arr, left, right) {
  let i = left;
  for (let j = left; j <= right; j++) {
    if (arr[j] <= arr[right]) {
      [arr[i], [arr[j]]] = [arr[j], [arr[i]]];
      i++;
    }
  }
  return i-1;
}

function quickSelect(arr, left, right, idxToFind) {
  if (left < right) {
    let partitionIdx = getPartitionIndex(arr, left, right);
    if (partitionIdx === idxToFind) {
      return arr[partitionIdx];
    }
    if (partitionIdx < idxToFind) {
      return quickSelect(arr, partitionIdx+1, right, idxToFind);
    }
    return quickSelect(arr, left, partitionIdx-1, idxToFind);
  }
}



var findKthLargest = function (nums, k) {
  const indexToFind = nums.length - k;

  return quickSelect(nums, 0, nums.length - 1, indexToFind);
};

module.exports = findKthLargest;