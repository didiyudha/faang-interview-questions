package main

var dictionary = map[string]string{
	"(": ")",
	"{": "}",
	"[": "]",
}

func IsValidParenthesis(s string) bool {

	stack := NewStack()

	for _, charByte := range s {
		char := string(charByte)

		_, ok := dictionary[char]
		if ok {
			stack.Push(char)
			continue
		}

		if stack.Cap() == 0 {
			return false
		}

		lastVal := stack.Pop()
		rightParen := dictionary[lastVal]

		if rightParen != char {
			return false
		}

	}

	return stack.Cap() == 0
}

func main() {

}