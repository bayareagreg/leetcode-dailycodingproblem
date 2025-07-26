package main

import "fmt"

// Trie represents the Trie data structure
type Trie struct {
	children []*Trie
}

// NewNode creates a new Trie node
func NewTrie() *Trie {
	return &Trie{
		children: make([]*Trie, 2),
	}
}

// Insert inserts a word into the Trie
func (t *Trie) Insert(n int) {
	curNode := t
	for i := 31; i >= 0; i-- {
		curBit := getIthBit(n, i)
		// if current bit is 0, it will go to children[1]
		// if current bit is 1, it will go to children[0]
		if curNode.children[curBit^1] == nil {
			curNode.children[curBit^1] = NewTrie()
		}
		curNode = curNode.children[curBit^1]
	}
}

func (t *Trie) FindMaxXOR(n int) int {
	max := 0
	curNode := t
	for i := 31; i >= 0; i-- {
		curBit := getIthBit(n, i)
		// if current bit is 1, then we want to go to children[1] (which represents 0)
		// if current bit is 0, then we want to go to children[0] (which represents 1)
		if curNode.children[curBit] != nil {
			curNode = curNode.children[curBit]
			// if exists, then we will have a "1" at the current index
			// in the result of maximum XOR
			max += (1 << i)
		} else {
			curNode = curNode.children[curBit^1]
		}
	}
	return max
}

func (t *Trie) _print(indent string) {
	if t == nil {
		fmt.Printf("%s'<nil>':\n", indent)
		return
	}
	for k, v := range t.children {
		fmt.Printf("%s'%v':\n", indent, k)
		v._print(indent + "  ")
	}
}

func (t *Trie) Print() {
	t._print("")
}

// retuns 0 or 1
func getIthBit(num int, n int) int {
	mask := 1 << n
	return (num & mask) >> n
}

func max_xor(nums []int) int {
	t := NewTrie()
	for i := range nums {
		t.Insert(nums[i])
	}

	t.Print()

	xor := 0
	for i := range nums {
		xor = max(xor, t.FindMaxXOR(nums[i]))
	}

	return xor
}

// O(n^2)
func max_xor2(arr []int) int {
	max := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			val := arr[i] ^ arr[j]
			if max < val {
				max = val
			}
		}
	}
	return max
}

func main() {
	nums := []int{3, 10, 5, 25, 2, 8}
	fmt.Printf("max XOR = %v\n", max_xor(nums))
}
