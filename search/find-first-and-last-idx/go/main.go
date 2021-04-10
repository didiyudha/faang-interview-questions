package main

import "math"

func binarySearch(nums []int, left, right, target int) int {
	for left <= right {
		middle := int(math.Floor((float64(left) + float64(right)) / 2))

		if nums[middle] == target {
			return middle
		}

		if nums[middle] < target {
			left = middle + 1
			continue
		}

		right = middle - 1
	}

	return -1
}

func searchRange(nums []int, target int) []int {

	if len(nums) == 0 {
		return []int{-1, -1}
	}

	var left int
	right := len(nums) - 1

	firstPos := binarySearch(nums, left, right, target)

	if firstPos == -1 {
		return []int{-1, -1}
	}

	startPos, endPos := firstPos, firstPos
	temp1, temp2 := -1, -1

	for startPos != -1 {
		temp1 = startPos
		startPos = binarySearch(nums, 0, startPos-1, target)
	}

	startPos = temp1

	for endPos != -1 {
		temp2 = endPos
		endPos = binarySearch(nums, endPos+1, len(nums)-1, target)
	}

	endPos = temp2

	return []int{startPos, endPos}

}

func main() {

}
