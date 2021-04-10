package main

import (
	"fmt"
	"math"
)
func binarySearch(nums []int, target int) (result int) {

	left := 0
	right := len(nums) - 1
	middle := getMiddlePoint(right, left)

	for left <= right {
		if nums[middle] == target {
			return nums[middle]
		}
		if nums[middle] < target {
			left = middle + 1
			middle = getMiddlePoint(right, left)
		}
		if nums[middle] > target {
			right = middle - 1
			middle = getMiddlePoint(right, left)
		}
	}

	return -1

}

func getMiddlePoint(right, left int) (middle int) {
	middle = int(math.Floor(float64(right+left) / float64(2)))
	return
}

func main() {

	nums := []int{1, 2, 3, 4, 5 , 6, 7, 8, 9}
	target := 100

	result := binarySearch(nums, target)
	fmt.Println(result)
}
