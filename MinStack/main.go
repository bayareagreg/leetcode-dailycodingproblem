package main

import "fmt"

type MinStack struct {
	arr  []int
	mins []int
}

func Constructor() MinStack {
	return MinStack{
		arr:  make([]int, 0),
		mins: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.arr = append(this.arr, val)
	if len(this.mins) == 0 {
		this.mins = append(this.mins, val)
	} else {
		this.mins = append(this.mins,
			min(val, this.mins[len(this.mins)-1]))
	}
}

func (this *MinStack) Pop() {
	this.arr = this.arr[:len(this.arr)-1]
	this.mins = this.mins[:len(this.mins)-1]
}

func (this *MinStack) Top() int {
	if len(this.arr) == 0 {
		panic("Stack is empty")
	}
	return this.arr[len(this.arr)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.mins) == 0 {
		panic("Stack is empty")
	}
	return this.mins[len(this.mins)-1]
}

func (this *MinStack) Print() {
	fmt.Printf("Stack: ")
	for i := range this.arr {
		fmt.Printf(" %v", this.arr[i])
	}
	fmt.Printf("\n. mins: %v", this.mins)
	fmt.Println()
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Printf("min: %v\n", obj.GetMin())
	obj.Pop()
	fmt.Printf("top: %v\n", obj.Top())
	fmt.Printf("min: %v\n", obj.GetMin())
}
