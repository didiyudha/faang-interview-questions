package main

func main() {

}

func TrapWater(height []int) int {

	var total int

	pLeft, pRight := 0, len(height)-1
	maxLeft, maxRight := 0, 0

	for pLeft < pRight {
		if height[pLeft] <= height[pRight] {
			// Move from left to right.
			if maxLeft < height[pLeft] {
				maxLeft = height[pLeft]
			} else {
				total += maxLeft - height[pLeft]
			}
			pLeft++
		} else {
			// Move from right to left.
			if maxRight < height[pRight] {
				maxRight = height[pRight]
			} else {
				total += maxRight - height[pRight]
			}
			pRight--
		}
	}

	return total
}
