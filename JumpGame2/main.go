package main

import (
	"fmt"
	"math"
)

func jump(nums []int) int {
	dp := make([]int, len(nums))
	dp[len(nums)-1] = 0
	for i := len(nums) - 2; i >= 0; i-- {
		minMoves := math.MaxInt32
		for j := i + 1; j < len(nums); j++ {
			if nums[i] >= (j - i) {
				minMoves = min(minMoves, dp[j]+1)
			}
		}
		dp[i] = minMoves
	}
	return dp[0]

}

func main() {
	nums := []int{2, 2, 0, 1, 4}
	//      []int{3, 2, INF, 1, 0}
	fmt.Println(jump(nums))
}
