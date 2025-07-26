package main

import "fmt"

func findDuplicate(nums []int) int {
	t := nums[0]
	h := nums[0]
	for {
		t = nums[t]
		h = nums[nums[h]]
		if t == h {
			break
		}
	}

	// Step 2: Find the entrance to the cycle (duplicate number)
	t = nums[0]
	for t != h {
		t = nums[t]
		h = nums[h]
	}
	return t
}

func main() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(nums))
}
