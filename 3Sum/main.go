// To execute Go code, please declare a func main() in a package "main"

package main

import (
	"fmt"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]] = i
	}

	m2 := make(map[[3]int]struct{})
	for i := range nums {
		for j := range nums {
			if i != j {
				k, ok := m[-(nums[i] + nums[j])]
				_, ok2 := m2[[3]int{nums[i], nums[j], nums[k]}]
				if k != i && k != j && ok && !ok2 {
					m2[[3]int{nums[i], nums[j], nums[k]}] = struct{}{}
					m2[[3]int{nums[i], nums[k], nums[j]}] = struct{}{}
					m2[[3]int{nums[j], nums[k], nums[i]}] = struct{}{}
					m2[[3]int{nums[j], nums[i], nums[k]}] = struct{}{}
					m2[[3]int{nums[k], nums[i], nums[j]}] = struct{}{}
					m2[[3]int{nums[k], nums[j], nums[i]}] = struct{}{}
					res = append(res, []int{nums[i], nums[j], nums[k]})
					//					delete(m, -(nums[i] + nums[j]))
				}
			}
		}
	}
	return res

}

func main() {
	fmt.Printf("%v\n", threeSum([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}))
}
