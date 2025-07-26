package main

import "fmt"

// Node represents a node in the Trie
type Node struct {
	children map[rune]*Node
	isEnd    bool
}

// Trie represents the Trie data structure
type Trie struct {
	root *Node
}

// NewNode creates a new Trie node
func NewNode() *Node {
	return &Node{
		children: make(map[rune]*Node),
		isEnd:    false,
	}
}

// NewTrie initializes a new Trie
func NewTrie() *Trie {
	return &Trie{
		root: NewNode(),
	}
}

// Insert inserts a word into the Trie
func (t *Trie) Insert(word string) {
	current := t.root
	for _, char := range word {
		if _, ok := current.children[char]; !ok {
			current.children[char] = NewNode()
		}
		current = current.children[char]
	}
	current.isEnd = true
}

func (t *Trie) Search(word string) bool {
	current := t.root
	for _, char := range word {
		if _, ok := current.children[char]; !ok {
			return false
		}
		current = current.children[char]
	}
	return current.isEnd
}

func (n *Node) _print(indent string) {
	if n.isEnd {
		fmt.Printf("%s##end##\n", indent)
	}
	for k, v := range n.children {
		fmt.Printf("%s'%c':\n", indent, k)
		v._print(indent + "  ")
	}
}

func (t *Trie) Print() {
	t.root._print("")
}

func main() {
	t := NewTrie()
	t.Insert("dog")
	t.Insert("dingo")
	t.Insert("zebra")
	t.Insert("cat")
	/*
		t.Insert("columnar")
		t.Insert("column")
	*/
	t.Print()

	fmt.Printf("Search = %t\n", t.Search("column"))
}
