function maxArea(arr) {
  let max = 0;

  for (let i = 0; i < arr.length; i++) {
    for (let j = i+1; j < arr.length; j++) {
      const height = Math.min(arr[i], arr[j]);
      const width = j - i;
      const area = countArea(height, width);
      max = Math.max(max, area);
    }
  }

  return max;
}

function optimalMaxArea(arr) {
  
  let max = 0;
  let p1 = 0;
  let p2 = arr.length-1;

  while(p1 < p2) {
    const height = Math.min(arr[p1], arr[p2]);
    const width = p2 - p1;
    const area = countArea(height, width);
    max = Math.max(max, area);

    if (arr[p1] < arr[p2]) {
      p1++;
    } else {
      p2--;
    }

  }

  return max;

}

function countArea(height, width) {
  return height * width;
}

module.exports = {
  maxArea,
  optimalMaxArea,
};