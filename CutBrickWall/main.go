package main

import (
	"fmt"
	"math"
)

type Pair struct {
	x int
	y int
}

func leastBricks(wall [][]int) int {
	numCols := 0
	edges := make(map[Pair]struct{})
	for i := range wall {
		x := 0
		for j := range wall[i] {
			x += wall[i][j]
			edges[Pair{x: x, y: i}] = struct{}{}
		}
		if i == 0 {
			numCols = x
		}
	}

	minCuts := math.MaxInt
	for i := 0; i < numCols; i++ {
		curCuts := 0
		for j := range wall {
			_, ok := edges[Pair{x: i, y: j}]
			if !ok {
				curCuts++
			}
		}
		minCuts = min(minCuts, curCuts)
	}
	return minCuts
}

func main() {
	/*
		wall1 := [][]int{
			{3, 5, 1, 1},
			{2, 3, 3, 2},
			{5, 5},
			{4, 4, 2},
			{1, 3, 3, 3},
			{1, 1, 6, 1, 1}}
		fmt.Printf("%v\n", leastBricks(wall1))
	*/

	wall2 := [][]int{
		{1},
		{1},
		{1},
	}
	fmt.Printf("%v\n", leastBricks(wall2))
}
