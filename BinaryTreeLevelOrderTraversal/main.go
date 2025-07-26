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

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	res := [][]int{}
	for len(queue) > 0 {
		curLevelCount := len(queue)
		curLevelArray := []int{}
		for i := range curLevelCount {
			node := queue[i]
			curLevelArray = append(curLevelArray, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, curLevelArray)
		queue = queue[curLevelCount:]
	}
	return res
}

func main() {
	root := &TreeNode{
		3,
		&TreeNode{9, nil, nil},
		&TreeNode{
			20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil},
		},
	}
	fmt.Printf("%v\n", levelOrder(root))
}
