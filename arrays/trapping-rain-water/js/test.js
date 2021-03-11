const { trapRainWater, optimalTrapRainWater } = require('./index');

describe('Bruteforce', () => {
  
  test('should return 8', () => {
    const heights = [0, 1, 0, 2, 1, 0, 3, 1, 0, 1, 2];
    const total = trapRainWater(heights);
    expect(total).toEqual(8);
  });

  test('should return 6', () => {
    const inp = [0,1,0,2,1,0,1,3,2,1,2,1];
    const total = trapRainWater(inp);
    expect(total).toEqual(6);
  });
  
});

describe('Optimal solution', () => {
  test('should return 8', () => {
    const heights = [0, 1, 0, 2, 1, 0, 3, 1, 0, 1, 2];
    const total = optimalTrapRainWater(heights);
    expect(total).toEqual(8);
  })

  test('should return 6', () => {
    const inp = [0,1,0,2,1,0,1,3,2,1,2,1];
    const total = optimalTrapRainWater(inp);
    expect(total).toEqual(6);
  });
  
});