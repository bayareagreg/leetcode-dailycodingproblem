package main

import (
	"fmt"
	"math"
)

// O(n+m)
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	merged := []int{}
	if len(nums1) == 0 {
		merged = nums2
	} else if len(nums2) == 0 {
		merged = nums1
	} else {
		i := 0
		j := 0
		for {
			if nums1[i] < nums2[j] {
				merged = append(merged, nums1[i])
				i++
			} else if nums1[i] > nums2[j] {
				merged = append(merged, nums2[j])
				j++
			} else {
				merged = append(merged, nums1[i])
				i++
				merged = append(merged, nums2[j])
				j++
			}

			if i >= len(nums1) {
				if j < len(nums2) {
					merged = append(merged, nums2[j:]...)
				}
				break
			} else if j >= len(nums2) {
				if i < len(nums1) {
					merged = append(merged, nums1[i:]...)
				}
				break
			}
		}
	}
	//fmt.Printf("%v\n", merged)
	mid := len(merged) / 2
	if len(merged)%2 == 1 {
		return float64(merged[mid])
	} else {
		return float64((float64(merged[mid-1]) + float64(merged[mid])) / 2)
	}
}

func seenElement(num int, seen [2]int) [2]int {
	if seen[0] == math.MinInt32 {
		seen[0] = num
	} else if seen[1] == math.MinInt32 {
		seen[1] = num
	} else {
		seen[0] = seen[1]
		seen[1] = num
	}
	return seen
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	even := true
	seen := [2]int{math.MinInt32, math.MinInt32}
	last := (len(nums1) + len(nums2)) / 2
	if (len(nums1)+len(nums2))%2 == 1 {
		even = false
	}
	i := 0
	j := 0
	for {
		if i >= len(nums1) {
			if j < len(nums2) {
				seen = seenElement(nums2[j], seen)
				j++
			} else {
				panic("7")
			}
		} else if j >= len(nums2) {
			if i < len(nums1) {
				seen = seenElement(nums1[i], seen)
				i++
			} else {
				panic("8")
			}
		} else if nums1[i] < nums2[j] {
			seen = seenElement(nums1[i], seen)
			i++
		} else {
			seen = seenElement(nums2[j], seen)
			j++
		}

		if i+j > last {
			break
		}
	}
	fmt.Printf("%v\n", seen)
	if even {
		return float64((float64(seen[0]) + float64(seen[1])) / 2)
	} else {
		if seen[1] == math.MinInt32 {
			return float64(seen[0])
		} else {
			return float64(seen[1])
		}
	}
}

func main() {
	//nums1 := []int{1, 3}
	//nums2 := []int{2}

	//nums1 := []int{1, 2}
	//nums2 := []int{3, 4}

	//nums1 := []int{1}
	//nums2 := []int{}

	nums1 := []int{2, 2, 4, 4}
	nums2 := []int{2, 2, 4, 4}

	fmt.Printf("%v\n", findMedianSortedArrays(nums1, nums2))
}
