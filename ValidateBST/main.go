package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, res)
	*res = append(*res, root.Val)
	inOrder(root.Right, res)
}

func isValidBST(root *TreeNode) bool {
	res := []int{}
	inOrder(root, &res)
	//	fmt.Printf("%v\n", res)
	for i := 1; i < len(res); i++ {
		if res[i] <= res[i-1] {
			return false
		}
	}
	return true
}

func main() {
	/*
		root := &TreeNode{
			2,
			&TreeNode{1, nil, nil},
			&TreeNode{3, nil, nil},
		}
	*/
	root := &TreeNode{
		2,
		&TreeNode{2, nil, nil},
		&TreeNode{2, nil, nil},
	}
	/*
			root := &TreeNode{
				5,
				&TreeNode{1, nil, nil},
				&TreeNode{
					4,
					&TreeNode{3, nil, nil},
					&TreeNode{6, nil, nil},
				},
			}
		root := &TreeNode{
			5,
			&TreeNode{4, nil, nil},
			&TreeNode{
				6,
				&TreeNode{3, nil, nil},
				&TreeNode{7, nil, nil},
			},
		}
	*/
	fmt.Printf("%t\n", isValidBST(root))
}
