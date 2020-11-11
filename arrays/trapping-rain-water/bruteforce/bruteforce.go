package bruteforce

func GetTotalWater(heights []int) int {
	var totalWater int
	for p := 0; p < len(heights); p++ {
		leftP := p
		rightP := p
		var maxLeft int
		var maxRight int

		for leftP >= 0 {
			maxLeft = Max(maxLeft, heights[leftP])
			leftP--
		}

		for rightP < len(heights) {
			maxRight = Max(maxRight, heights[rightP])
			rightP++
		}

		currentWater := Min(maxLeft, maxRight) - heights[p]

		if currentWater >= 0 {
			totalWater += currentWater
		}

	}

	return totalWater
}

func Max(nums ...int)  (max int){
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return
}

func Min(nums ...int)  (min int){
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