package main

import (
	"fmt"
)

// https://en.wikipedia.org/wiki/Levenshtein_distance
func minDistance(word1 string, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)

	prevRow := make([]int, len2+1)
	currentRow := make([]int, len2+1)

	for p := range len2 + 1 {
		prevRow[p] = p
	}

	for s1 := 1; s1 <= len1; s1++ {
		currentRow[0] = s1
		for s2 := 1; s2 < len(prevRow); s2++ {
			if word1[s1-1] == word2[s2-1] {
				// The last chars of the two substrings are the same, no operation needed
				currentRow[s2] = prevRow[s2-1]
			} else {
				// Dealing with the char difference (plus 1) and choosing the optimal way
				currentRow[s2] = min(
					currentRow[s2-1], // add char OR
					prevRow[s2-1],    // substitution OR
					prevRow[s2]) + 1  // deletion
			}
		}
		//	fmt.Printf("prevRow: %v, currentRow: %v\n", prevRow, currentRow)
		prevRow, currentRow = currentRow, prevRow
	}
	return prevRow[len2]
}

func main() {
	//fmt.Printf("%v\n", minDistance("horse", "ros"))
	//fmt.Printf("%v\n", minDistance("intention", "execution"))
	//fmt.Printf("%v\n", minDistance("sea", "eat"))
	fmt.Printf("%v\n", minDistance("dinitrophenylhydrazine", "benzalphenylhydrazone"))
}
