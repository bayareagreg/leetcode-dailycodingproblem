package main

import "fmt"

// find the right row using binary chop
// or return -1 if there is no applicable row
func verticalBinaryChop(matrix [][]int, target int) int {
	// a couple of corner cases
	if target < matrix[0][0] {
		return -1
	}

	m := len(matrix)
	n := len(matrix[0])

	if target > matrix[m-1][0] {
		return m - 1
	}

	lo, hi := 0, m-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if matrix[mid][0] == target {
			return mid
		} else if matrix[mid][0] > target {
			hi = mid - 1
		} else if matrix[mid][n-1] >= target {
			return mid
		} else {
			lo = mid + 1
		}
	}
	return -1
}

func horizontalBinaryChop(array []int, target int) bool {
	lo, hi := 0, len(array)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if array[mid] == target {
			return true
		} else if array[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	row := verticalBinaryChop(matrix, target)
	if row == -1 {
		return false
	}
	return horizontalBinaryChop(matrix[row], target)
}

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	fmt.Printf("%v\n", searchMatrix(matrix, 13))
}
