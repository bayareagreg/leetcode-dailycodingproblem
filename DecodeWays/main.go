package main

import "fmt"

func numDecodings(s string) int {
	dp := make([]int, len(s))
	if s[0] != '0' {
		dp[0] = 1
	}
	if len(s) == 1 {
		return dp[0]
	}

	for i := 1; i < len(s); i++ {
		// if single char is valid
		if s[i] > '0' && s[i] <= '9' {
			dp[i] = dp[i-1]
		}

		// if two chars valid
		if s[i-1] == '1' || s[i-1] == '2' && s[i] < '7' {
			if i-2 >= 0 {
				dp[i] += dp[i-2]
			} else {
				dp[i] += 1
			}
		}
	}

	return dp[len(s)-1]
}

func main() {
	fmt.Println(numDecodings("12"))
}
