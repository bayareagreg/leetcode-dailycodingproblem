package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	negative := false
	if x < 0 {
		x = -x
		negative = true
	}
	digits := []int{}
	for i, n := 0, 10; i < 10; i, n = i+1, 10*n {
		d := (x % n) / (n / 10)
		x -= d
		digits = append(digits, d)
	}
	//	fmt.Printf("%v\n", digits)
	// find the most significant digit
	msd := 0
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 0 {
			msd = i
			break
		}
	}
	n := int(math.Pow(float64(10), float64(msd)))
	x = 0
	for i := 0; i <= msd; i++ {
		x += n * digits[i]
		if x > math.MaxInt32 {
			return 0
		}
		n = n / 10
	}
	if negative {
		return -x
	} else {
		return x
	}
}

func main() {
	fmt.Printf("%v\n", reverse(1534))
}
