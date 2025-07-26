// To execute Go code, please declare a func main() in a package "main"

package main

import (
	"fmt"
)

func tower_of_hanoi(n int, a, b, c rune) {
	if n >= 1 {
		tower_of_hanoi(n-1, a, c, b)
		fmt.Printf("Move %c to %c\n", a, c)
		tower_of_hanoi(n-1, b, a, c)
	}
}

func main() {
	tower_of_hanoi(4, '1', '2', '3')
}
