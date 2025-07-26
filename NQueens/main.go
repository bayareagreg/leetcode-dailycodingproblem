package main

import (
	"fmt"
)

func abs(val int) int {
	if val < 0 {
		return -val
	} else {
		return val
	}
}

func is_valid(board []int) bool {
	current_queen_row, current_queen_col := len(board)-1, board[len(board)-1]

	// check if any queens can attack the last queen
	for row := range len(board) - 1 {
		col := board[row]
		diff := abs(current_queen_col - col)
		if diff == 0 || diff == current_queen_row-row {
			return false
		}
	}
	return true
}

func convertBoardToStringArray(board []int) []string {
	res := []string{}
	for i := range board {
		row := ""
		for j := range board {
			if j == board[i] {
				row += "Q"
			} else {
				row += "."
			}
		}
		res = append(res, row)
	}
	return res
}

func solveNQueensHelper(n int, board []int, res *[][]string) {
	if n == len(board) {
		*res = append(*res, convertBoardToStringArray(board))
	}

	for col := range n {
		board = append(board, col)
		if is_valid(board) {
			solveNQueensHelper(n, board, res)
		}
		board = board[0 : len(board)-1]
	}
}

func solveNQueens(n int) [][]string {
	board := []int{}
	result := [][]string{}
	solveNQueensHelper(n, board, &result)
	return result
}

func main() {
	fmt.Printf("%v\n", solveNQueens(5))
}
