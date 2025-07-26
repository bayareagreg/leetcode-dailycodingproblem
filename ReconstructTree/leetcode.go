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

func printHelper(t *TreeNode, indent int) {
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
		printHelper(t.Left, indent+2)
		printHelper(t.Right, indent+2)
	}
}

func (t *TreeNode) Print() {
	printHelper(t, 0)
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	if len(preorder) == 1 && len(inorder) == 1 {
		return &TreeNode{preorder[0], nil, nil}
	}

	root_index := indexOf(inorder, preorder[0])
	if root_index < 0 {
		panic("no root index")
	}

	left_pre := []int{}
	right_pre := []int{}
	left_in := []int{}
	right_in := []int{}

	num_left := root_index
	num_right := len(preorder) - num_left - 1

	if num_left > 0 {
		left_pre = preorder[1 : num_left+1]
		left_in = inorder[0:num_left]
	}
	if num_right > 0 {
		right_pre = preorder[root_index+1:]
		right_in = inorder[root_index+1:]
	}

	return &TreeNode{
		preorder[0],
		buildTree(left_pre, left_in),   // left
		buildTree(right_pre, right_in)} // right
}

func main() {
	preoder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	t := buildTree(preoder, inorder)

	t.Print()

}
