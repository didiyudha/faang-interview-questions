package bruteforce

func TwoSum(nums []int, target int) []int {
	for p1 := 0; p1 < len(nums); p1++ {
		numToFind := target - nums[p1]
		for p2 := p1 + 1; p2 < len(nums); p2++ {
			if nums[p2] == numToFind {
				return []int{p1, p2}
			}
		}
	}
	return nil
}
