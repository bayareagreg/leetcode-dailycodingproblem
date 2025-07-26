package main

import "fmt"

func checkColumns(board [][]byte) bool {
	// check columns
	for j := 0; j < len(board); j++ {
		ar := make([]int, 9)
		for i := range board {
			c := board[i][j]
			if '1' <= c && c <= '9' {
				ix := c - '1'
				if ar[ix] != 0 {
					return false
				} else {
					ar[ix] = 1
				}
			} else if c == '.' {
				continue
			} else {
				return false
			}
		}
	}
	return true
}

func checkRows(board [][]byte) bool {
	// check rows
	for i := 0; i < len(board); i++ {
		ar := make([]int, 9)
		for j := range board[i] {
			c := board[i][j]
			if '1' <= c && c <= '9' {
				ix := c - '1'
				if ar[ix] != 0 {
					return false
				} else {
					ar[ix] = 1
				}
			} else if c == '.' {
				continue
			} else {
				return false
			}
		}
	}
	return true
}

func checkNoRepetitions(board [][]byte) bool {
	ar := make([]int, 9)
	for i := 0; i < len(board); i++ {
		for j := range board[i] {
			c := board[i][j]
			if '1' <= c && c <= '9' {
				ix := c - '1'
				if ar[ix] != 0 {
					return false
				} else {
					ar[ix] = 1
				}
			} else if c == '.' {
				continue
			}
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	if !checkRows(board) || !checkColumns(board) {
		return false
	}
	for k := 0; k < 9; k = k + 3 {
		for j := 0; j < 9; j = j + 3 {
			slice := [][]byte{}
			for i := k; i < k+3; i++ {
				slice = append(slice, board[i][j:j+3])
			}
			print(slice)
			if !checkNoRepetitions(slice) {
				return false
			}
		}
	}
	return true
}

func print(slice [][]byte) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++ {
			fmt.Printf("%c", slice[i][j])
		}
		fmt.Println()
	}
	fmt.Println("----")
}

func main() {
	_ = [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	input2 := [][]byte{
		{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
		{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
		{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
		{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
		{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
		{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
		{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
	}
	fmt.Printf("%t\n", isValidSudoku(input2))
}
