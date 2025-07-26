package main

import (
	"fmt"
	"math"
)

// two pointers
func maxArea(height []int) int {
	max := math.MinInt32
	start, end := 0, len(height)-1
	for start < end {
		h := min(height[start], height[end])
		w := end - start
		if h*w > max {
			max = h * w
		}
		if height[start] > height[end] {
			end--
		} else if height[start] < height[end] {
			start++
		} else {
			end--
			start++
		}
	}
	return max
}

// brute force
func maxArea2(height []int) int {
	max := math.MinInt32
	for i := range height {
		for j := i; j < len(height); j++ {
			h := min(height[i], height[j])
			w := j - i
			if h*w > max {
				max = h * w
			}
		}
	}
	return max
}

func main() {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Printf("%v\n", maxArea(input))
}
