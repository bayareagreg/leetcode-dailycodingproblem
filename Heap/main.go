package main

import (
	"container/heap"
	"fmt"
)

// IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Peek() int          { return h[0] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	/*
		h := &IntHeap{2, 1, 5}
		// heapify
		heap.Init(h)
		heap.Push(h, 3)
		fmt.Printf("minimum: %d\n", (*h)[0]) // Output: minimum: 1
		for h.Len() > 0 {
			fmt.Printf("%d ", heap.Pop(h)) // Output: 1 2 3 5
		}
		fmt.Println()
	*/
	h := &IntHeap{}
	heap.Push(h, 2)
	heap.Push(h, 3)
	heap.Push(h, 8)
	heap.Push(h, 6)
	fmt.Printf("heap = %v\n", h)
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h)) // Output: 1 2 3 5
	}
}
