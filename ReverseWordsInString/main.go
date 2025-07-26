package main

import (
	"fmt"
	"slices"
	"strings"
)

func reverseWords(s string) string {
	res := ""
	inWord := false
	curWord := ""

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if !inWord {
				continue
			} else {
				inWord = false
				if res != "" {
					res = res + " "
				}
				res += curWord
				curWord = ""
			}
		} else {
			if !inWord {
				inWord = true
			}
			curWord = string(s[i]) + curWord
		}
	}
	if inWord {
		if res != "" {
			res = res + " "
		}
		res += curWord
	}
	return res
}

func reverseWords2(s string) string {
	fields := strings.Fields(strings.TrimSpace(s))
	slices.Reverse(fields)
	return strings.Join(fields, " ")
}

func main() {
	//input := "the sky is blue"
	//input := "  hello world  "
	input := "a good   example"

	fmt.Printf("'%v'\n", reverseWords(input))
}
