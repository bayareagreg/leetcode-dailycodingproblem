package main

import (
	"fmt"
)

// TrieNode represents a node in the Trie
type TrieNode struct {
	letters map[rune]*TrieNode
	total   int
}

type PrefixMapSum struct {
	root *TrieNode
	m    map[string]int
}

// NewTrieNode creates a new Trie node
func NewTrieNode() *TrieNode {
	return &TrieNode{
		letters: make(map[rune]*TrieNode),
		total:   0,
	}
}

// NewTrie initializes a new Trie
func NewPrefixMapSum() *PrefixMapSum {
	return &PrefixMapSum{
		root: NewTrieNode(),
		m:    make(map[string]int),
	}
}

// Insert inserts a word into the Trie
func (t *PrefixMapSum) Insert(word string, val int) {
	if val2, ok := t.m[word]; ok {
		val -= val2
	}

	trie := t.root
	for _, char := range word {
		if _, ok := trie.letters[char]; !ok {
			trie.letters[char] = NewTrieNode()
		}
		trie = trie.letters[char]
		trie.total += val
	}
}

func (n *TrieNode) _print(indent string) {
	for k, v := range n.letters {
		fmt.Printf("%s'%c': %v\n", indent, k, v.total)
		v._print(indent + "  ")
	}
}

func (t *PrefixMapSum) Print() {
	t.root._print("")
}

// must be O(1)
func (p *PrefixMapSum) Sum(prefix string) int {
	current := p.root
	for _, char := range prefix {
		if _, ok := current.letters[char]; !ok {
			return -1
		}
		current = current.letters[char]
	}
	return current.total
}

func main() {
	p := NewPrefixMapSum()
	p.Insert("columnar", 3)
	p.Print()
	fmt.Printf("%v\n", p.Sum("col"))
	p.Insert("column", 2)
	p.Print()
	fmt.Printf("%v\n", p.Sum("col"))
	fmt.Printf("%v\n", p.Sum("column"))
}
