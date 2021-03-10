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

module.exports = {
  trapRainWater,
};