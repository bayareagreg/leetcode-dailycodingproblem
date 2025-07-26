package main

import "fmt"

func canJump(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[len(nums)-1] = true
	for i := len(nums) - 2; i >= 0; i-- {
		canJump := false
		for j := i + 1; j < len(nums); j++ {
			if dp[j] && nums[i] >= (j-i) {
				canJump = true
				break
			}
		}
		dp[i] = canJump
	}
	return dp[0]
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	//nums := []int{3, 2, 1, 0, 4}
	//               i  j
	fmt.Println(canJump(nums))
}
