package main

import "fmt"

func setZeroes(matrix [][]int) {
	n := len(matrix)
	k := len(matrix[0])
	zero_rows := make([]int, n)
	zero_columns := make([]int, k)
	for i := range n {
		for j := range k {
			if matrix[i][j] == 0 {
				zero_rows[i] = 1
				zero_columns[j] = 1
			}
		}
	}
	for i := range n {
		if zero_rows[i] == 1 {
			for j := range k {
				matrix[i][j] = 0
			}
		}
	}

	for j := range k {
		if zero_columns[j] == 1 {
			for i := range n {
				matrix[i][j] = 0
			}
		}
	}
}

func main() {
	input := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1}}

	setZeroes(input)
	fmt.Printf("%v\n", input)
}
