package main

import "fmt"

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	} else if len(nums) == 2 {
		return [][]int{{nums[0], nums[1]}, {nums[1], nums[0]}}
	} else {
		res := [][]int{}
		for _, p := range permute(nums[1:]) {
			for i := range nums {
				// need to insert nums[0] into i-th position
				x := []int{}
				x = append(x, p[0:i]...)
				x = append(x, nums[0])
				x = append(x, p[i:]...)
				res = append(res, x)
			}
		}
		return res
	}
}

func main() {
	input := []int{1, 2, 3}
	fmt.Printf("%v\n", permute(input))

}
