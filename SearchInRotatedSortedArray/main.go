package main

import "fmt"

func search2(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}

// input := []int{4, 5, 6, 7, 0, 1, 2}
// pivot at k: num[k] > num[k+1], num[k+1] < num[k+2]
func findPivot(nums []int) int {
	lo, hi := 0, len(nums)-1
	if hi == lo || nums[lo] < nums[hi] {
		// there is no pivot
		return -1
	}

	for lo <= hi {
		if lo == hi-1 {
			if nums[lo] > nums[hi] {
				return hi
			} else {
				break
			}
		} else {
			mid := lo + (hi-lo)/2
			// is mid the pivot?
			if nums[mid] < nums[mid-1] && nums[mid] < nums[mid+1] {
				return mid
			} else if nums[lo] < nums[mid] {
				// pivot is in this half
				// mid..hi
				lo = mid
			} else {
				// pivot is in this half
				// lo .. mid
				hi = mid
			}
		}
	}
	return -1
}

func search(nums []int, target int) int {
	p := findPivot(nums)
	if p == -1 {
		return search2(nums, target)
	} else {
		nums2 := nums[p:]
		nums2 = append(nums2, nums[0:p]...)
		n := search2(nums2, target)
		if n == -1 {
			return -1
		}
		return (p + n) % len(nums)
	}
}

func main() {
	//input := []int{4, 5, 6, 7, 0, 1, 2}
	input := []int{3, 1}
	fmt.Printf("%v\n", search(input, 3))
	// original
	// {0, 1, 2, 4, 5, 6, 7}
	// pivot = 5
	// {6, 7, 0, 1, 2, 4, 5}
	//input := []int{6, 7, 8, 5}
	//input := []int{6, 7, 8, 5}

	//fmt.Printf("Pivot: %v\n", findPivot(input))

}
