package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func NewTreeNode(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{
		val:   val,
		left:  left,
		right: right,
	}
}

func printHelper(t *TreeNode, indent int) {
	if t == nil {
		return
	}
	for _ = range indent {
		fmt.Printf(" ")
	}
	fmt.Printf("%c", t.val)
	fmt.Println()
	printHelper(t.left, indent+2)
	printHelper(t.right, indent+2)
}

func (t *TreeNode) Print() {
	printHelper(t, 0)
}

func helper(t *TreeNode) (int, bool, int) {
	if t == nil {
		return 0, false, 0
	}
	if t.left == nil && t.right == nil {
		return t.val, true, 1
	}
	if t.left != nil && t.right != nil {
		val, is_unival, count := helper(t.left)
		val2, is_unival2, count2 := helper(t.right)
		count3 := count + count2
		if t.val == val && t.val == val2 && is_unival && is_unival2 {
			count3++
			return t.val, true, count3
		} else {
			return t.val, false, count3
		}
	}
	if t.left != nil && t.right == nil {
		val, is_unival, count := helper(t.left)
		if t.val == val && is_unival {
			count++
			return t.val, true, count
		} else {
			return t.val, false, count
		}
	}
	val, is_unival, count := helper(t.right)
	if t.val == val && is_unival {
		count++
		return t.val, true, count
	} else {
		return t.val, false, count
	}
}

func countUnivalTrees(t *TreeNode) int {
	_, _, count := helper(t)
	return count
}

func main() {
	/*
			t := NewTreeNode(0,
				NewTreeNode(1, nil, nil),
				NewTreeNode(0,
					NewTreeNode(
						1,
						NewTreeNode(1, nil, nil),
						NewTreeNode(1, nil, nil)),
					NewTreeNode(0, nil, nil)))
			fmt.Printf("count = %v\n", countUnivalTrees(t))
		t :=
			NewTreeNode(
				'a',
				NewTreeNode('a', nil, nil),
				NewTreeNode(
					'a',
					NewTreeNode('a', nil, nil),
					NewTreeNode(
						'a',
						nil,
						NewTreeNode('b', nil, nil))))
		fmt.Printf("count = %v\n", countUnivalTrees(t))
	*/
	t :=
		NewTreeNode(
			'a',
			NewTreeNode('c', nil, nil),
			NewTreeNode(
				'b',
				NewTreeNode('b', nil, nil),
				NewTreeNode(
					'b',
					nil,
					NewTreeNode('b', nil, nil))))
	//fmt.Printf("count = %v\n", countUnivalTrees(t))
	t.Print()
}
