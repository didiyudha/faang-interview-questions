package main

import (
	"fmt"
	"strings"
)

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

func longestSubstrOptimal(s string) (total int) {
	var left int
	m := make(map[string]int)
	arr := strings.Split(s, "")
	fmt.Println("arr: ", arr)

	for right := 0; right < len(arr); right++ {
		currentChar := arr[right]
		prevSeen, ok := m[currentChar]
		if ok && prevSeen >= left {
			left = prevSeen + 1
		}
		m[currentChar] = right
		currentLen := right - left + 1
		fmt.Println("left: ", left)
		fmt.Println("right: ", right)
		fmt.Println("current len: ", currentLen)
		total = Max(total, currentLen)
	}

	return total
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
