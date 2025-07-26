package main

import "fmt"

/**
 * Definition for a Node.
 */
type Node struct {
	Val       int
	Neighbors []*Node
}

func createNodes(node *Node, m map[int]*Node) {
	if node == nil {
		return
	}
	_, ok := m[node.Val]
	if !ok {
		m[node.Val] = &Node{node.Val, make([]*Node, 0)}
	}
	for _, n := range node.Neighbors {
		_, ok := m[n.Val] // like visited check
		if !ok {
			createNodes(n, m)
		}
	}
}

func createNeighbors(node *Node, m map[int]*Node) {
	if node == nil {
		return
	}
	for _, n := range node.Neighbors {
		m[node.Val].Neighbors = append(m[node.Val].Neighbors, m[n.Val])
	}

	for _, n := range node.Neighbors {
		if len(m[n.Val].Neighbors) == 0 {
			createNeighbors(n, m)
		}
	}
}

func printGraph(m map[int]*Node) {
	for i := 1; i <= 4; i++ {
		node := m[i]
		fmt.Print("[")
		for _, n := range node.Neighbors {
			fmt.Printf("%v,", n.Val)
		}
		fmt.Print("],")
	}
	fmt.Println()
}

func cloneGraph(node *Node) *Node {
	m := make(map[int]*Node)
	createNodes(node, m)
	createNeighbors(node, m)
	printGraph(m)
	return m[1]
}

func main() {
	n1 := &Node{1, nil}
	n2 := &Node{2, nil}
	n3 := &Node{3, nil}
	n4 := &Node{4, nil}
	n1.Neighbors = []*Node{n2, n4}
	n2.Neighbors = []*Node{n1, n3}
	n3.Neighbors = []*Node{n2, n4}
	n4.Neighbors = []*Node{n1, n3}
	cloneGraph(n1)
}
