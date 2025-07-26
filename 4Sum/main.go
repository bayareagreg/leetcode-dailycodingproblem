package main

import (
	"fmt"
	"sort"
)

// all combinations of a, b, c, d
func combos(a, b, c, d int) [][4]int {
	res := [][4]int{}
	input := []int{a, b, c, d}
	for i1, v1 := range input {
		for i2, v2 := range input {
			for i3, v3 := range input {
				for i4, v4 := range input {
					if i1 != i2 && i1 != i3 && i2 != i3 && i1 != i4 && i2 != i4 && i3 != i4 {
						res = append(res, [4]int{v1, v2, v3, v4})
					}
				}
			}
		}
	}
	//fmt.Printf("combos for %v, %v, %v, %v: %v\n", a, b, c, d, res)
	return res
}

func fourSum2(nums []int, target int) [][]int {
	res := [][]int{}
	m := make(map[int][]int)
	for i := range nums {
		arr, ok := m[nums[i]]
		if !ok {
			arr = []int{}
		}
		m[nums[i]] = append(arr, i)
	}

	m2 := make(map[[4]int]struct{})
	for a := range nums {
		for b := range nums {
			if a != b {
				for c := range nums {
					if b != c && a != c {
						if arr, ok := m[target-(nums[a]+nums[b]+nums[c])]; ok {
							for _, d := range arr {
								if d != a && d != b && d != c {
									if _, ok2 := m2[[4]int{nums[a], nums[b], nums[c], nums[d]}]; !ok2 {
										for _, c := range combos(nums[a], nums[b], nums[c], nums[d]) {
											m2[c] = struct{}{}
										}
										res = append(res, []int{nums[a], nums[b], nums[c], nums[d]})
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return res
}

func fourSum(nums []int, target int) [][]int {
	var result [][]int
	length := len(nums)
	sort.Ints(nums)
	for i := 0; i < length-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		if nums[i]+nums[length-1]+nums[length-2]+nums[length-3] < target {
			continue
		}
		for j := i + 1; j < length-2; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			if nums[i]+nums[j]+nums[length-1]+nums[length-2] < target {
				continue
			}
			head, end := j+1, length-1
			for head < end {
				tempt := nums[i] + nums[j] + nums[head] + nums[end]
				if tempt == target {
					var lengthArray []int
					lengthArray = append(lengthArray, nums[i])
					lengthArray = append(lengthArray, nums[j])
					lengthArray = append(lengthArray, nums[head])
					lengthArray = append(lengthArray, nums[end])
					result = append(result, lengthArray)
					head++
					for head < end && nums[head] == nums[head-1] {
						head++
					}
				} else if tempt > target {
					end--
				} else {
					head++
				}
			}
		}
	}
	return result
}

func main() {
	//input := []int{1, 0, -1, 0, -2, 2}
	//fmt.Printf("%v\n", fourSum(input, 0))

	//input := []int{2, 2, 2, 2, 2}
	//fmt.Printf("%v\n", fourSum(input, 8))

	//input := []int{-2, -1, -1, 1, 1, 2, 2}
	//fmt.Printf("%v\n", fourSum(input, 0))

	input := []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
	fmt.Printf("%v\n", fourSum(input, 8))

}
