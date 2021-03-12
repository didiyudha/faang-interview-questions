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

func backspaceCompareOptimal(s, t string) bool {
	arrS, arrT := strings.Split(s, ""), strings.Split(t, "")
	p1, p2 := len(arrS)-1, len(arrT)-1

	for p1 >= 0 || p2 >= 0 {
		if arrS[p1] == "#" || arrT[p2] == "#" {
			if arrS[p1] == "#" {
				backCount := 2
				for backCount > 0 {
					p1--
					backCount--

					for p1 >= 0 && arrS[p1] == "#" {
						backCount += 2
					}
				}
			}
			if arrT[p2] == "#" {
				backCount := 2
				for backCount > 0 {
					p2--
					backCount--

					for p2 >= 0 && arrT[p2] == "#" {
						backCount += 2
					}
				}
			}
		} else {

			if arrS[p1] != arrT[p2] {
				return false
			}

			p1--
			p2--

		}
	}

	return true
}
