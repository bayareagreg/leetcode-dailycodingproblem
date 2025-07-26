package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	longestPalindrome := s[0:1]

	is_palindrome := func(s string) bool {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}

	for i := range s {
		for j := i + len(longestPalindrome) + 1; j <= len(s); j++ {
			// this is the new string to check
			// start:  i
			// end: i+len(longestPalindrome) + 1
			check := s[i:j]
			// check if palindrome
			if is_palindrome(check) {
				longestPalindrome = check
			}
		}
	}
	return longestPalindrome
}

func main() {
	fmt.Printf("%v\n", longestPalindrome("bb"))

	// fmt.Printf("%v\n", longestPalindrome("babad"))
	// fmt.Printf("%v\n", longestPalindrome("cbbd"))
}
