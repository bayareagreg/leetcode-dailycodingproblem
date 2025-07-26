package main

import (
	"fmt"
	"math"
)

// KADANE
func maxSubArray(nums []int) int {
	bestSum := math.MinInt32
	curSum := 0
	for _, x := range nums {
		curSum = max(x, curSum+x)
		bestSum = max(bestSum, curSum)
	}
	return bestSum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Printf("%v\n", maxSubArray(nums))
}
