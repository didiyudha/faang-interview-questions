function kthLargest(arr, k) {
  arr.sort((a, b) => a-b);
  
  const idx = arr.length - k;
  return arr[idx];
}

module.exports = kthLargest;