package main

func main() {

}

func IsPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if string(s[left]) != string(s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}
