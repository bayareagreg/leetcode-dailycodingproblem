package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

// O(n^3logn)
func generateRegularNumbers(n int) {
	out := make(chan int)

	var generator func(int, chan<- int)
	generator = func(n int, out chan<- int) {
		defer close(out)

		twos := []int{}
		for i := range n {
			twos = append(twos, int(math.Pow(float64(2), float64(i))))
		}
		threes := []int{}
		for i := range n {
			threes = append(threes, int(math.Pow(float64(3), float64(i))))
		}
		fives := []int{}
		for i := range n {
			fives = append(fives, int(math.Pow(float64(5), float64(i))))
		}

		all := []int{}
		for i := range twos {
			for j := range threes {
				for k := range fives {
					all = append(all, twos[i]*threes[j]*fives[k])
				}
			}
		}

		sort.Ints(all)

		for i := range all[0:n] {
			out <- all[i]
		}
	}

	go generator(n, out)
	for x := range out {
		fmt.Printf("%v\n", x)
	}
}

// IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

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

func generateRegularNumbers2(n int) {
	out := make(chan int)

	var generator func(int, chan<- int)
	generator = func(n int, out chan<- int) {
		defer close(out)

		h := &IntHeap{1}
		count := 0
		last := 0
		for count < n {
			val := heap.Pop(h).(int)
			if val > last {
				out <- val
				last = val
				count++
				heap.Push(h, val*2)
				heap.Push(h, val*3)
				heap.Push(h, val*5)
			}
		}
	}

	go generator(n, out)
	for x := range out {
		fmt.Printf("%v\n", x)
	}
}

func main() {
	generateRegularNumbers2(10)
}
