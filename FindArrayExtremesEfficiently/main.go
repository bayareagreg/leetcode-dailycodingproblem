package main

import "fmt"

func findArrayExtremes(arr []int) (int, int) {
	if len(arr) == 1 {
		return arr[0], arr[0]
	} else if len(arr) == 2 {
		if arr[0] < arr[1] {
			return arr[0], arr[1]
		} else {
			return arr[1], arr[0]
		}
	} else {
		n := len(arr) / 2
		min1, max1 := findArrayExtremes(arr[0:n])
		min2, max2 := findArrayExtremes(arr[n+1:])
		return min(min1, min2), max(max1, max2)

	}
}

func main() {
	input := []int{4, 2, 7, 5, -1, 3, 6}
	min, max := findArrayExtremes(input)
	fmt.Printf("min: %v, max: %v\n", min, max)
}
