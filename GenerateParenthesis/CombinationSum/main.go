package main

import (
	"fmt"
	"sort"
)

func InsertIntoSortedSlice(slice []int, value int) []int {
	// Find the insertion index using sort.Search
	idx := sort.Search(len(slice), func(i int) bool {
		return slice[i] >= value
	})

	// Create a new slice with one extra capacity
	newSlice := make([]int, 0, len(slice)+1)

	// Append elements before the insertion point
	newSlice = append(newSlice, slice[:idx]...)

	// Append the new value
	newSlice = append(newSlice, value)

	// Append elements after the insertion point
	newSlice = append(newSlice, slice[idx:]...)

	return newSlice
}

func AreSortedSlicesEqual(a, b []int) bool {
	// If lengths are different, slices cannot be equal.
	if len(a) != len(b) {
		return false
	}

	// Compare elements at corresponding indices.
	for i := range a {
		if a[i] != b[i] {
			return false // Found a differing element
		}
	}

	return true // All elements are equal
}

func FindThisSlice(slice []int, in [][]int) bool {
	for _, s := range in {
		if AreSortedSlicesEqual(s, slice) {
			return true
		}
	}
	return false
}

func AddSliceToResult(slice []int, result [][]int) [][]int {
	if !FindThisSlice(slice, result) {
		return append(result, slice)
	} else {
		return result
	}
}

func AddSlicesToResult(slices [][]int, result [][]int) [][]int {
	for _, s := range slices {
		result = AddSliceToResult(s, result)
	}
	return result
}

func tryThis(candidates []int, path []int, target int) [][]int {
	if target < candidates[0] {
		// nope, impossible
		return nil
	} else {
		res := [][]int{}
		for i := range candidates {
			if candidates[i] == target {
				newpath := InsertIntoSortedSlice(path, target)
				res = AddSliceToResult(newpath, res)
			} else if candidates[i] < target {
				newpath := InsertIntoSortedSlice(path, candidates[i])
				res2 := tryThis(candidates, newpath, target-candidates[i])
				if res2 != nil {
					res = AddSlicesToResult(res2, res)
				}
			}
		}
		return res
	}
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	for i := range candidates {
		path := []int{candidates[i]}
		if candidates[i] == target {
			res = AddSliceToResult(path, res)
		} else {
			res2 := tryThis(candidates, path, target-candidates[i])
			if res2 != nil {
				res = AddSlicesToResult(res2, res)
			}
		}
	}
	return res
}

func main() {
	//candidates := []int{2, 3, 5}
	//candidates := []int{2}
	candidates := []int{2, 3, 6, 7}
	fmt.Printf("%v\n", combinationSum(candidates, 7))
}
