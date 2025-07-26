package main

import "fmt"

func generateParenthesis(n int) []string {
	combinations := []string{}

	var backtrack func(currentCombo string, opened int, closed int)

	backtrack = func(currentCombo string, opened int, closed int) {
		// base case if len is n*2 that means we got combination
		// wanted
		if len(currentCombo) == n*2 {
			combinations = append(combinations, currentCombo)
			return
		}
		// in every recursion there are 2 possible ways close or open
		// but we have to make sure that max number opening and
		// closing parenthesis is n and to close you have to have enough opening parenthesis
		if opened < n {
			backtrack(currentCombo+"(", opened+1, closed)
		}
		if opened > closed && closed < n {
			backtrack(currentCombo+")", opened, closed+1)
		}
	}
	backtrack("", 0, 0)

	return combinations
}

func main() {
	fmt.Printf("%v\n", generateParenthesis(4))
}
