const quickSort = require('./index');

test('Should sort the array', function() {
  const arr = [5,3,1,6,4,2];
  const left = 0;
  const right = arr.length-1;

  quickSort(arr, left, right);

  for (let i = 0; i < arr.length; i++) {
    expect(arr[i]).toEqual(i+1);
  }

})