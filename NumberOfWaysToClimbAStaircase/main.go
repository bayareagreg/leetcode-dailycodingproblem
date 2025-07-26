package main

import (
	"fmt"
)

func numberOfWaysToClimb(n int) [][]int {

	if n == 1 {
		return [][]int{{1}}
	} else if n == 2 {
		return [][]int{{1, 1}, {2}}
	} else {
		steps := make(map[string]struct{})
		result := [][]int{}
		prev := numberOfWaysToClimb(n - 1)
		for _, num := range prev {
			num1 := append(num, 1)
			num2 := append([]int{1}, num...)
			str1 := fmt.Sprintf("%v", num1)
			if _, ok := steps[str1]; !ok {
				steps[str1] = struct{}{}
				result = append(result, num1)
			}
			str2 := fmt.Sprintf("%v", num2)
			if _, ok := steps[str2]; !ok {
				steps[str2] = struct{}{}
				result = append(result, num2)
			}
		}
		prev2 := numberOfWaysToClimb(n - 2)
		for _, num := range prev2 {
			num1 := append(num, 2)
			num2 := append([]int{2}, num...)
			str1 := fmt.Sprintf("%v", num1)
			if _, ok := steps[str1]; !ok {
				steps[str1] = struct{}{}
				result = append(result, num1)
			}
			str2 := fmt.Sprintf("%v", num2)
			if _, ok := steps[str2]; !ok {
				steps[str2] = struct{}{}
				result = append(result, num2)
			}
		}
		return result
	}
}

func numberOfWaysToClimb2(n int) int {
	cache := make(map[int]int)
	for i := 0; i < n+1; i++ {
		cache[i] = 0
	}
	cache[0] = 1
	for i := 1; i < n+1; i++ {
		cache[i] += cache[i-1] + cache[i-2]
	}
	return cache[n]
}

func main() {
	num := numberOfWaysToClimb2(8)
	fmt.Printf("%v ways\n", num)
	//fmt.Printf("%v ways: %v\n", len(num), num)
}
