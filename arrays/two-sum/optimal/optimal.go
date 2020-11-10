package optimal

func TwoSum(nums []int, target int) []int {
	numMaps := make(map[int]int)
	for p, n := range nums{
		idx, ok := numMaps[n]
		if ok {
			return []int{idx, p}
		}
		numToFind := target - n
		numMaps[numToFind] = p
	}
	return nil
}
