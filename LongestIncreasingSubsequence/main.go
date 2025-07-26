package main

import "fmt"

func lengthOfLIS(nums []int) int {
	//binary search with tail array
	tail := []int{}
	for _, num := range nums {
		idx := binarySearch(tail, num)
		if idx == len(tail) {
			tail = append(tail, num)
		} else {
			tail[idx] = num
		}
	}
	return len(tail)
}

// returns an index of a element in an array
func binarySearch(tail []int, target int) int {
	n := len(tail)
	left, right := 0, n
	for left < right {
		mid := left + (right-left)/2
		if tail[mid] < target {
			left = mid + 1
		} else if tail[mid] > target {
			right = mid
		} else {
			return mid
		}
	}
	return left
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	//nums := []int{0, 1, 0, 3, 2, 3}
	//nums := []int{7, 7, 7, 7, 7, 7, 7}
	//nums := []int{4, 10, 4, 3, 8, 9}
	fmt.Println(lengthOfLIS(nums))
}
