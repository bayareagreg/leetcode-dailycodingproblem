package main

import "fmt"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  left,
		Right: right,
	}
}

func (t *TreeNode) _print(indent int) {
	for range indent {
		fmt.Printf(" ")
	}
	if t == nil {
		fmt.Printf("<nil>\n")
		return
	}
	fmt.Printf("%v", t.Val)
	fmt.Println()
	if t.Left != nil || t.Right != nil {
		t.Left._print(indent + 2)
		t.Right._print(indent + 2)
	}
}

func (t *TreeNode) Print() {
	t._print(0)
}

func make_trees(low, high int) []*TreeNode {
	trees := make([]*TreeNode, 0)
	if low > high {
		trees = append(trees, nil)
		return trees
	}
	for i := low; i < high+1; i++ {
		left := make_trees(low, i-1)
		right := make_trees(i+1, high)
		for j := range left {
			for k := range right {
				trees = append(trees, NewTreeNode(i, left[j], right[k]))
			}
		}
	}
	return trees
}

func generateTrees(n int) []*TreeNode {
	return make_trees(1, n)
}

func main() {
	tt := generateTrees(5)
	for i := range tt {
		tt[i].Print()
		fmt.Println("-----------------")
	}
	fmt.Printf("Total trees: %v\n", len(tt))
}
