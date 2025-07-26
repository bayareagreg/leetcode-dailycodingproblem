package main

import (
	"container/heap"
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

func NewNode(data int) *Node {
	return &Node{
		data:  data,
		left:  nil,
		right: nil,
	}
}

type NodeHeap []*Node

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].data < h[j].data }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h NodeHeap) Peek() *Node        { return h[0] }

func (h *NodeHeap) Push(x any) {
	*h = append(*h, x.(*Node))
}

func (h *NodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *NodeHeap) Print() {
	fmt.Printf("NodeHeap: [")
	for i := range *h {
		fmt.Printf("{%v,%v,%v} ", (*h)[i].left, (*h)[i].data, (*h)[i].right)
	}
	fmt.Printf("]\n")
}

func buildFrequences(s string) map[byte]int {
	freq := make(map[byte]int)
	for i := range s {
		if _, ok := freq[s[i]]; !ok {
			freq[s[i]] = 0
		}
		freq[s[i]] = freq[s[i]] + 1
	}
	return freq
}

func preOrder(root *Node, ans *[]string, curr string) {
	if root == nil {
		return
	}

	// Leaf node represents a character.
	if root.left == nil && root.right == nil {
		*ans = append(*ans, curr)
		return
	}

	preOrder(root.left, ans, curr+string('0'))
	preOrder(root.right, ans, curr+string('1'))
}

func huffmanCodes(s string) []string {
	freq := buildFrequences(s)

	// Min heap for node class.
	pq := &NodeHeap{}
	for _, f := range freq {
		heap.Push(pq, NewNode(f))
	}
	pq.Print()

	// Construct huffman tree.
	for pq.Len() >= 2 {

		// Left node
		l := heap.Pop(pq).(*Node)

		// Right node
		r := heap.Pop(pq).(*Node)

		newNode := NewNode(l.data + r.data)
		newNode.left = l
		newNode.right = r

		heap.Push(pq, newNode)
	}

	pq.Print()

	root := heap.Pop(pq).(*Node)
	ans := []string{}
	preOrder(root, &ans, "")
	return ans
}

func main() {
	input := "aaaaabbbbbbbbbccccccccccccdddddddddddddeeeeeeeeeeeeeeeefffffffffffffffffffffffffffffffffffffffffffff"
	fmt.Printf("%v\n", huffmanCodes(input))
}
