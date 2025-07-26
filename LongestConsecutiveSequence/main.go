package main

import (
	"fmt"
	"sort"
)

// O(nlogn)
func longestConsecutive2(nums []int) int {
	res := 1
	sort.Ints(nums)
	cur := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i]-1 {
			cur++
			res = max(res, cur)
		} else {
			cur = 1
		}
	}
	return res
}

func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})
	for _, n := range nums {
		m[n] = struct{}{}
	}
	maxlen := 0
	for k := range m {
		if _, ok := m[k-1]; ok {
			// k cannot be the start
			continue
		}
		cur := 1
		for {
			k++
			if _, ok := m[k]; !ok {
				break
			}
			cur++
		}
		maxlen = max(maxlen, cur)
	}
	return maxlen
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}
