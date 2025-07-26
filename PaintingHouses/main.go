package main

import (
	"fmt"
	"math"
)

type Cost struct {
	colors []int
	cost   int
}

func permutations(n, k int) [][]int {
	res := [][]int{}
	if n == 1 {
		for i := range k {
			res = append(res, []int{i})
		}
	} else {
		x := permutations(n-1, k)
		for _, p := range x {
			for i := range k {
				if i != p[len(p)-1] {
					row := append(p, i)
					res = append(res, row)
				}
			}
		}
	}
	return res
}

// brute force - O(n^k)
func build_houses2(matrix [][]int) int {
	costs := []Cost{}

	n := len(matrix)
	k := len(matrix[0])

	fmt.Printf("Houses: %v, Colors: %v\n", n, k)

	minCost := math.MaxInt32
	for _, p := range permutations(n, k) {
		cost := 0
		for i := range p {
			cost += matrix[i][p[i]]
		}
		c := Cost{
			colors: p,
			cost:   cost,
		}
		costs = append(costs, c)
		minCost = min(minCost, cost)
	}

	fmt.Printf("Costs: %v\n", costs)
	return minCost
}

func build_houses(matrix [][]int) int {
	n := len(matrix)
	k := len(matrix[0])

	prevCosts := make([]int, k)
	currCosts := make([]int, k)

	// Initialize for the first house
	for j := range k {
		prevCosts[j] = matrix[0][j]
	}

	// Iterate through the remaining houses
	for i := 1; i < n; i++ {
		for j := range k {
			minPrev := math.MaxInt32
			for l := range k {
				if j != l {
					minPrev = min(minPrev, prevCosts[l])
				}
			}
			currCosts[j] = matrix[i][j] + minPrev
		}
		// Update prevCosts for the next iteration
		for j := range k {
			prevCosts[j] = currCosts[j]
		}
	}

	// Find the minimum cost from the last house's painting options
	minTotalCost := math.MaxInt32
	for j := range k {
		minTotalCost = min(minTotalCost, prevCosts[j])
	}
	return minTotalCost
}

func main() {
	//n := 3
	//k := 3
	matrix := [][]int{
		{1, 2, 10},
		{6, 3, 9},
		{4, 5, 6},
	}

	fmt.Printf("%v\n", build_houses(matrix))
}
