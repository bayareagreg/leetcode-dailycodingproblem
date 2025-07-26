package main

import (
	"errors"
	"fmt"
	"math"
)

func myAtoi(s string) int {
	negative := false
	sign_read := false
	digits := []int{}
	for _, c := range s {
		if c == ' ' {
			if sign_read || len(digits) > 0 {
				break
			}
			continue
		} else if c == '+' {
			if sign_read || len(digits) > 0 {
				break
			}
			sign_read = true
			continue
		} else if c == '-' {
			if sign_read || len(digits) > 0 {
				break
			}
			negative = true
			sign_read = true
			continue
		} else {
			digit, err := charToDigit(c)
			if err != nil {
				break
			}
			digits = append(digits, digit)
		}
	}

	val := 0
	for i, n := len(digits)-1, 1; i >= 0; i, n = i-1, n*10 {
		val += n * digits[i]
		if n > math.MaxInt32 {
			for ; i >= 0; i-- {
				if digits[i] != 0 {
					if negative {
						val = -math.MinInt32
					} else {
						val = math.MaxInt32
					}
					break
				}
			}
			break
		}
		if negative {
			if -val < math.MinInt32 {
				val = -math.MinInt32
				break
			}
		} else if val > math.MaxInt32 {
			val = math.MaxInt32
			break
		}
	}
	if negative {
		return -val
	} else {
		return val
	}
}

func charToDigit(c rune) (int, error) {
	switch c {
	case '0':
		return 0, nil
	case '1':
		return 1, nil
	case '2':
		return 2, nil
	case '3':
		return 3, nil
	case '4':
		return 4, nil
	case '5':
		return 5, nil
	case '6':
		return 6, nil
	case '7':
		return 7, nil
	case '8':
		return 8, nil
	case '9':
		return 9, nil
	default:
		return 0, errors.New("not a digit")

	}
}

func main() {
	fmt.Printf("%v\n", myAtoi("-91283472332"))
}
