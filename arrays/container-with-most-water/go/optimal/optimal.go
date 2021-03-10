package optimal

func MaxArea(height []int) int {

	var maxArea int

	p1 := 0
	p2 := len(height) - 1

	for p1 < p2 {
		h := findHeight(height[p1], height[p2])
		w := p2 - p1
		currentArea := findMaxArea(h, w)
		if maxArea < currentArea {
			maxArea = currentArea
		}
		if height[p1] < height[p2] {
			p1++
		} else {
			p2--
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
