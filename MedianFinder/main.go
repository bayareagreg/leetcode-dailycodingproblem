package main

import (
	"container/heap"
	"fmt"
)

type Heap struct {
	Values   []int
	LessFunc func(int, int) bool
}

func (h *Heap) Less(i, j int) bool { return h.LessFunc(h.Values[i], h.Values[j]) }
func (h *Heap) Swap(i, j int)      { h.Values[i], h.Values[j] = h.Values[j], h.Values[i] }
func (h *Heap) Len() int           { return len(h.Values) }
func (h *Heap) Peek() int          { return h.Values[0] }
func (h *Heap) Pop() (v interface{}) {
	h.Values, v = h.Values[:h.Len()-1], h.Values[h.Len()-1]
	return v
}
func (h *Heap) Push(v interface{}) { h.Values = append(h.Values, v.(int)) }

func NewHeap(less func(int, int) bool) *Heap {
	return &Heap{LessFunc: less}
}

type MedianFinder struct {
	smallHeap *Heap
	largeHeap *Heap
}

func Constructor() MedianFinder {
	return MedianFinder{
		smallHeap: NewHeap(func(a, b int) bool {
			return a > b
		}),
		largeHeap: NewHeap(func(a, b int) bool {
			return a < b
		}),
	}
}

func (mf *MedianFinder) AddNum(num int) {
	if (mf.smallHeap.Len()+mf.largeHeap.Len())%2 == 0 {
		heap.Push(mf.largeHeap, num)
		heap.Push(mf.smallHeap, heap.Pop(mf.largeHeap))
	} else {
		heap.Push(mf.smallHeap, num)
		heap.Push(mf.largeHeap, heap.Pop(mf.smallHeap))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if (mf.smallHeap.Len()+mf.largeHeap.Len())%2 == 0 {
		return (float64(mf.smallHeap.Peek()) + float64(mf.largeHeap.Peek())) / 2
	}
	return float64(mf.smallHeap.Peek())
}

func (this *MedianFinder) Print() {
	fmt.Printf("MedianFinder {\n")
	fmt.Printf("  minHeap: %v\n", this.smallHeap.Values)
	fmt.Printf("  maxHeap: %v\n", this.largeHeap.Values)
	fmt.Printf("}\n")
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
func main() {
	//stream := []int{2, 1, 5, 7, 2, 0, 5}
	stream := []int{1, 2, 3, 7, 4}

	obj := Constructor()
	for i := range stream {
		obj.AddNum(stream[i])
		obj.Print()
		fmt.Printf("%v\n", obj.FindMedian())
	}
}
