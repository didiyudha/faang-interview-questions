function trapRainWater(arr) {

  let totalArea = 0;

  for (let i = 0; i < arr.length; i++) {

    let pLeft = i;
    let pRight = i;
    let maxLeft = 0;
    let maxRight = 0;

    while(pLeft >= 0) {
      maxLeft = Math.max(maxLeft, arr[pLeft]);
      pLeft--;
    }

    while(pRight < arr.length) {
      maxRight = Math.max(maxRight, arr[pRight]);
      pRight++;
    }

    const area = Math.min(maxRight, maxLeft) - arr[i];
    if (area > 0) {
      totalArea += area;
    }

  }

  return totalArea;
  
}

function optimalTrapRainWater(height) {
  let totalArea = 0;
  let maxLeft = 0, maxRight = 0;
  let left = 0, right = height.length-1;

  while (left < right) {
    if (height[left] <= height[right]) {
      if (height[left] > maxLeft) {
        maxLeft = height[left];
      } else {
        totalArea += maxLeft-height[left];
      }
      left++
    } else {
      if (height[right] > maxRight) {
        maxRight = height[right];
      } else {
        totalArea += maxRight - height[right];
      }
      right--;
    }
  }

  return totalArea;
}

module.exports = {
  trapRainWater,
  optimalTrapRainWater,
};