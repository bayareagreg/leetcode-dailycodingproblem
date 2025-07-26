package main

import "fmt"

func nextLetter(board [][]byte, i, j int) [][2]int {
	res := [][2]int{}
	// up to 4 coordinates
	// 1) up
	if i > 0 {
		res = append(res, [2]int{i - 1, j})
	}
	// 2) right
	if j < len(board[0])-1 {
		res = append(res, [2]int{i, j + 1})
	}
	// 3) down
	if i < len(board)-1 {
		res = append(res, [2]int{i + 1, j})
	}
	// 4) right
	if j > 0 {
		res = append(res, [2]int{i, j - 1})
	}
	return res
}

func search(board [][]byte, i, j int, word string, visited map[[2]int]struct{}) bool {
	//	fmt.Printf("+search(%v, %v, %v)\n", i, j, word)

	if word[0] == board[i][j] {
		if len(word) == 1 {
			return true
		}
		visited[[2]int{i, j}] = struct{}{}
		for _, n := range nextLetter(board, i, j) {
			if _, ok := visited[[2]int{n[0], n[1]}]; !ok {
				if search(board, n[0], n[1], word[1:], visited) {
					return true
				}
			}
		}
		delete(visited, [2]int{i, j})
	}
	return false
}

// I bet this has something to do with DFS
func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	for i := range m {
		for j := range n {
			if board[i][j] == word[0] {
				visited := make(map[[2]int]struct{})
				if search(board, i, j, word, visited) {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	input := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'}}

	fmt.Printf("%t\n", exist(input, "ABCEFSADEESE"))
	/*
			input := [][]byte{{'a'}}
			fmt.Printf("%t\n", exist(input, "a"))

		input := [][]byte{
			{'C', 'A', 'A'},
			{'A', 'A', 'A'},
			{'B', 'C', 'D'}}
		fmt.Printf("%t\n", exist(input, "AAB"))
	*/

}
