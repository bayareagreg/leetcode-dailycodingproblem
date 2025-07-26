package main

import "fmt"

// red: 0
// white: 1
// blue: 2
func sortColors(nums []int) {
	r, w, b := 0, 0, 0
	for _, c := range nums {
		switch c {
		case 0:
			r++
		case 1:
			w++
		case 2:
			b++
		}
	}
	for i := range nums {
		if r > 0 {
			nums[i] = 0
			r--
		} else if w > 0 {
			nums[i] = 1
			w--
		} else {
			nums[i] = 2
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Printf("%v\n", nums)
}
