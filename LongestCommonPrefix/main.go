package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	res := ""
	first := strs[0]
	for i := 0; i < len(first); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) {
				return res
			}
			if first[i] != strs[j][i] {
				return res
			}
		}
		res = res + string(first[i])
	}
	return res
}

func main() {
	fmt.Printf("%v\n", longestCommonPrefix([]string{"ab", "a"}))
}
