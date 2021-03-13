package main

func main() {

}

func IsAlmostPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		leftCharacter := string(s[left])
		rightCharacter := string(s[right])

		if leftCharacter != rightCharacter {
			return IsSubPalindrome(s, left+1, right) || IsSubPalindrome(s, left, right-1)
		}

		left++
		right--
	}

	return true

}

func IsSubPalindrome(s string, left, right int) bool {
	for left < right {
		leftCharacter := string(s[left])
		rightCharacter := string(s[right])

		if leftCharacter != rightCharacter {
			return false
		}
		left++
		right--
	}

	return true
}
