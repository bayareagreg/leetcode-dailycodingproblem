package main

import "fmt"

func rotate(nums []int, k int) {
	k = k % len(nums)
	result := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	for i := range nums {
		nums[i] = result[i]
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	//nums := []int{-1, -100, 3, 99}
	//nums := []int{1, 2, 3}
	k := 3

	rotate(nums, k)
	fmt.Println(nums)
}
