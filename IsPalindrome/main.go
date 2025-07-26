package main

import "fmt"

func isPalindrome(x int) bool {
	reverseX := 0
	for num := x; num > 0; {
		rem := num % 10
		reverseX = reverseX*10 + rem
		num = num / 10
	}
	return x == reverseX
}

func main() {
	fmt.Printf("%t\n", isPalindrome(1000001))
}
