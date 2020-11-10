package bruteforce

func MaxArea(height []int) int {
	var maxArea int
	for p1 := 0; p1 < len(height); p1++ {
		for p2 := p1+1; p2 < len(height); p2++ {
			height := findHeight(height[p1], height[p2])
			width := p2 - p1
			area := findMaxArea(height, width)
			if maxArea < area {
				maxArea = area
			}
		}
	}
	return maxArea
}

func findHeight(h1, h2 int) int {
	if h1 == h2 {
		return h1
	}
	if h1 < h2 {
		return h1
	}
	return h2
}

func findMaxArea(height, width int) int {
	return height * width
}
