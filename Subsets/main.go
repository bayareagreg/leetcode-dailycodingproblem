package main

import "fmt"

func subsets(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{}, {nums[0]}}
	} else {
		prev := subsets(nums[0 : len(nums)-1])
		res := [][]int{}
		res = append(res, prev...)
		for _, p := range prev {
			x := append([]int{}, p...)
			x = append(x, nums[len(nums)-1])
			res = append(res, x)
		}
		return res
	}
}

func main() {
	//nums := []int{1, 2, 3}
	nums := []int{9, 0, 3, 5, 7}
	fmt.Printf("%v\n", subsets(nums))
}
