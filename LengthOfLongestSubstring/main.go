package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	maxLen := 0

	for i := range s {
		for {
			// 1. generate string starting at i of maxLen+1
			if i+maxLen+1 > len(s) {
				break
			}
			sub := s[i : i+maxLen+1]
			// check if valid
			isValid := true
			m := make(map[byte]struct{})
			for j := range sub {
				if _, ok := m[sub[j]]; ok {
					isValid = false
					break
				}
				m[sub[j]] = struct{}{}
			}
			if isValid {
				maxLen++
			} else {
				break
			}
		}
		// if not continue main loop
	}
	return maxLen
}

func main() {
	//	fmt.Printf("%v\n", lengthOfLongestSubstring("abcabcbb"))
	fmt.Printf("%v\n", lengthOfLongestSubstring("a"))
}
