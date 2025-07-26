package main

import "fmt"

type TreeNode2 struct {
	val   int
	left  *TreeNode2
	right *TreeNode2
}

func NewTreeNode(val int, left *TreeNode2, right *TreeNode2) *TreeNode2 {
	return &TreeNode2{
		val:   val,
		left:  left,
		right: right,
	}
}

func printHelper2(t *TreeNode2, indent int) {
	for range indent {
		fmt.Printf(" ")
	}
	if t == nil {
		fmt.Printf("<nil>\n")
		return
	}
	fmt.Printf("%c", t.val)
	fmt.Println()
	if t.left != nil || t.right != nil {
		printHelper2(t.left, indent+2)
		printHelper2(t.right, indent+2)
	}
}

func (t *TreeNode2) Print2() {
	printHelper2(t, 0)
}

func indexOf(arr []int, val int) int {
	for i := range arr {
		if val == arr[i] {
			return i
		}
	}
	return -1
}

func reconstruct(pre []int, in []int) *TreeNode2 {
	if len(pre) == 0 || len(in) == 0 {
		return nil
	}
	if len(pre) == 1 && len(in) == 1 {
		return NewTreeNode(pre[0], nil, nil)
	}

	root_index := indexOf(in, pre[0])
	if root_index < 0 {
		panic("no root index")
	}

	left_pre := []int{}
	right_pre := []int{}
	left_in := []int{}
	right_in := []int{}

	num_left := root_index
	num_right := len(pre) - num_left - 1

	if num_left > 0 {
		left_pre = pre[1 : num_left+1]
		left_in = in[0:num_left]
	}
	if num_right > 0 {
		right_pre = pre[root_index+1:]
		right_in = in[root_index+1:]
	}

	return NewTreeNode(
		pre[0],
		reconstruct(left_pre, left_in),   // left
		reconstruct(right_pre, right_in)) // right
}

func main2() {
	pre := []int{'a', 'b', 'd', 'e', 'c', 'f', 'g'}
	in := []int{'d', 'b', 'e', 'a', 'f', 'c', 'g'}

	//pre := []int{'a', 'b', 'c'}
	//in := []int{'a', 'b', 'c'}

	//pre := []int{'a', 'b', 'c'}
	//in := []int{'c', 'b', 'a'}

	t := reconstruct(pre, in)
	t.Print2()
}
