package main

import (
	"fmt"
	"math"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func divide_slow(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	quotient := 0
	divid := abs(dividend)
	divis := abs(divisor)
	for ; divid >= divis; quotient++ {
		divid -= divis
	}

	negative := (dividend < 0) != (divisor < 0)
	if negative {
		return -quotient
	}

	return quotient
}

func divide_fast(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}

	negative := false
	if (dividend < 0) != (divisor < 0) {
		negative = true
	}

	dividend, divisor = abs(dividend), abs(divisor)
	if dividend < divisor {
		return 0
	}

	if divisor == 1 {
		if negative {
			return -dividend
		} else {
			return min(dividend, math.MaxInt32)
		}
	}

	// i don't understand this loop
	res := 0
	for dividend >= divisor {
		power := 1
		val := divisor
		for val+val <= dividend {
			val += val
			power += power
		}

		dividend -= val
		res += power
	}

	if negative {
		res = -res
	}

	if res > math.MaxInt32 {
		return math.MaxInt32
	}

	if res < math.MinInt32 {
		return math.MinInt32
	}

	return res
}

func main() {
	fmt.Printf("%v\n", divide_slow(10, 3))
	//fmt.Printf("%v\n", divide(-2147483648, 2))
}
