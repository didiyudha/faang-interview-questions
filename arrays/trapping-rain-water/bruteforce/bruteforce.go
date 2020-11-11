package bruteforce

func GetTotalWater(heights []int) int {
	var totalWater int

	for p, h := range heights{

		var maxLeft int
		var maxRight int

		leftP := p
		rightP := p

		for leftP >= 0 {
			maxLeft = Max(maxLeft, heights[leftP])
			leftP--
		}

		for rightP < len(heights) {
			maxRight = Max(maxRight, heights[rightP])
			rightP++
		}

		currentWaterArea := Min(maxLeft, maxRight) - h
		if currentWaterArea > 0 {
			totalWater += currentWaterArea
		}
	}

	return totalWater
}

func Max(nums ...int) (max int) {
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return
}

func Min(nums ...int) (min int) {
	for i, n := range nums {
		if i == 0 {
			min = n
			continue
		}
		if n < min {
			min = n
		}
	}
	return
}
