package main

import "fmt"

type Graph struct {
	m map[int][]int
}

func traverse(g *Graph, curr int, result map[int]int) int {
	descendants := 0
	for _, child := range g.m[curr] {
		num_nodes := traverse(g, child, result)
		result[child] += num_nodes - 1
		descendants += num_nodes
	}
	return descendants + 1
}

func maxEdgesRemoved(g *Graph) int {
	// DFS traversal to populate map which stores number of descendants
	// per node
	result := make(map[int]int)
	traverse(g, 1, result)
	//fmt.Printf("result: %v\n", result)
	count := 0
	for _, v := range result {
		if v%2 == 1 {
			count++
		}
	}
	return count
}

func main() {
	g := &Graph{
		m: make(map[int][]int),
	}
	g.m[1] = []int{2, 3}
	g.m[2] = []int{}
	g.m[3] = []int{4, 5}
	g.m[4] = []int{6, 7, 8}
	g.m[5] = []int{}
	g.m[6] = []int{}
	g.m[7] = []int{}
	g.m[8] = []int{}

	fmt.Printf("%v\n", maxEdgesRemoved(g))
}
