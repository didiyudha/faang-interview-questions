function maxArea(arr) {
  let max = 0;

  for (let i = 0; i < arr.length; i++) {
    for (let j = i+1; j < arr.length; j++) {
      
      const height = Math.min(arr[i], arr[j]);
      const width = j - i;
      const area = countArea(height, width);
      
      if (max < area) {
        max = area;
      }

    }
  }

  return max;
}

function countArea(height, width) {
  return height * width;
}

module.exports = {
  maxArea,
};