package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hello")
}

func backspaceCompare(s, t string) bool {
	return build(s) == build(t)
}

func build(str string) string {
	output := make([]string, 0, len(str))

	for _, s := range str {
		character := string(s)
		if character != "#" {
			output = append(output, character)
			continue
		}
		if len(output) > 0 {
			output = output[:len(output)-1]
		}
	}

	return strings.Join(output, "")
}
