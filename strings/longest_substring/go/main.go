package main

import "strings"

func main() {

}

func longestSubstr(s string) int {

	var longestSubstr int
	arr := strings.Split(s, "")

	for p1 := 0; p1 < len(arr); p1++ {

		m := make(map[string]bool)
		var currentSubstrLen int

		for p2 := p1; p2 < len(arr); p2++ {
			letter := arr[p2]
			_, ok := m[letter]
			if ok {
				break
			}
			m[letter] = true
			currentSubstrLen++
			longestSubstr = Max(longestSubstr, currentSubstrLen)
		}
	}

	return longestSubstr
}

func Max(numbers ...int) int {
	if len(numbers) == 0 {
		return -1
	}
	max := numbers[0]
	for _, n := range numbers {
		if max < n {
			max = n
		}
	}
	return max
}
