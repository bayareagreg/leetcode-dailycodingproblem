package main

import "fmt"

// binary chop
func findTarget(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (high + low) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func searchRange(nums []int, target int) []int {
	i := findTarget(nums, target)
	if i == -1 {
		return []int{-1, -1}
	} else {
		// go backwards or forwards
		low, high := i-1, i+1
		for {
			if low < 0 || nums[low] != target {
				low++
				break
			}
			low--
		}
		for {
			if high >= len(nums) || nums[high] != target {
				high--
				break
			}
			high++
		}
		return []int{low, high}
	}
}

func main() {
	input := []int{5, 7, 7, 8, 8, 10}
	fmt.Printf("%v\n", searchRange(input, 8))
}
