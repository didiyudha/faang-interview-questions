package main

import "strings"

func main() {

}

func MinimumBracket(s string) string {

	arr := strings.Split(s, "")
	stack := make([]int, 0, len(s))

	for i, v := range s {

		c := string(v)

		if c == "(" {
			stack = append(stack, i)
			continue
		}

		if c == ")" && len(stack) > 0 {
			stack = stack[:len(stack)-1]
			continue
		}

		if c == ")" {
			arr[i] = ""
		}
	}

	for _, idx := range stack {
		arr[idx] = ""
	}

	return strings.Join(arr, "")

}
