package main

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	if len(digits) == 1 {
		return digitToLetters(digits[0])
	} else {
		result := []string{}
		for _, l := range digitToLetters(digits[0]) {
			for _, c := range letterCombinations(digits[1:]) {
				result = append(result, l+c)
			}
		}
		return result
	}
}

func digitToLetters(digit byte) []string {
	switch digit {
	case '2':
		return []string{"a", "b", "c"}
	case '3':
		return []string{"d", "e", "f"}
	case '4':
		return []string{"g", "h", "i"}
	case '5':
		return []string{"j", "k", "l"}
	case '6':
		return []string{"m", "n", "o"}
	case '7':
		return []string{"p", "q", "r", "s"}
	case '8':
		return []string{"t", "u", "v"}
	case '9':
		return []string{"w", "x", "y", "z"}
	default:
		return []string{}
	}
}

func main() {
	fmt.Printf("%v\n", letterCombinations("23"))
}
