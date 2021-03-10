const { maxArea, optimalMaxArea } = require('./index');

describe('Bruteforce', () => {
  test('should return 32', () => {
    const heightsArray = [4,8,1,2,3,9];
    const max = maxArea(heightsArray);
    expect(max).toEqual(32);
  });
  
  test('should return 28', () => {
    const heightsArray = [7, 1, 2, 3, 9];
    const max = maxArea(heightsArray);
    expect(max).toEqual(28);
  });
  
  test('should return 15', () => {
    const heightsArray = [5, 2, 8, 7];
    const max = maxArea(heightsArray);
    expect(max).toEqual(15);
  });
});

describe('Optimal solution', () => {
  test('should return 32', () => {
    const heightsArray = [4,8,1,2,3,9];
    const max = optimalMaxArea(heightsArray);
    expect(max).toEqual(32);
  });
  
  test('should return 28', () => {
    const heightsArray = [7, 1, 2, 3, 9];
    const max = optimalMaxArea(heightsArray);
    expect(max).toEqual(28);
  });
  
  test('should return 15', () => {
    const heightsArray = [5, 2, 8, 7];
    const max = optimalMaxArea(heightsArray);
    expect(max).toEqual(15);
  });
});
