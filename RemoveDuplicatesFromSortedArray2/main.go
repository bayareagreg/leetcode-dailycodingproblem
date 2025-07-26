package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	j := 2                           // writer
	for i := 2; i < len(nums); i++ { // reader
		if nums[j-1] != nums[i] || nums[j-1] != nums[j-2] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}

func main() {
	//input := []int{1, 1, 1, 2, 2, 3}
	input := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	output := removeDuplicates(input)
	fmt.Printf("%v: %v\n", output, input)
}
