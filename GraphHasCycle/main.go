package main

import "fmt"

type Graph map[string][]string

// traverse all edges starting with a given vertex.
// if, for some vertex we find one of its neighbors has already been visited => cycle (return true)
// othewrwise return false
func search(g Graph, vertex string, visited map[string]struct{}, parent string, indent string) bool {
	fmt.Printf("%v+search(vertex=%s, visited=%v, parent=%v)\n", indent, vertex, visited, parent)
	visited[vertex] = struct{}{}
	for neighborIx := range g[vertex] {
		neighbor := g[vertex][neighborIx]
		if _, ok := visited[neighbor]; !ok {
			if search(g, neighbor, visited, vertex, indent+"  ") {
				return true
			}
		} else if parent != neighbor {
			return true
		}
	}
	return false
}

func HasCycle(g Graph) bool {
	visited := make(map[string]struct{})

	for k := range g {
		if _, ok := visited[k]; !ok {
			if search(g, k, visited, "", "") {
				return true
			}
		}
	}
	return false
}

func main() {
	g := make(Graph)
	g["a"] = []string{"b"}
	g["b"] = []string{"a", "c", "d"}
	g["c"] = []string{"b", "d"}
	g["d"] = []string{"c", "b"}

	fmt.Printf("%t\n", HasCycle(g))
}
