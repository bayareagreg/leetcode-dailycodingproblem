package main

import "fmt"

func swap(matrix [][]int, i1, j1, i2, j2 int) {
	matrix[i1][j1], matrix[i2][j2] = matrix[i2][j2], matrix[i1][j1]
}

func transpose(matrix [][]int) {
	for i := 1; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			swap(matrix, i, j, j, i)
		}
	}
}

func reverse(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix)/2; j++ {
			swap(matrix, i, j, i, len(matrix)-1-j)
		}
	}
}

func rotate(matrix [][]int) {
	transpose(matrix)
	reverse(matrix)
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(matrix)
	fmt.Printf("%v\n", matrix)

}
