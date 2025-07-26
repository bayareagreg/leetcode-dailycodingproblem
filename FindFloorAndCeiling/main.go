package main

import (
	"fmt"
)

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

func (t *TreeNode) _print(indent int) {
	for range indent {
		fmt.Printf(" ")
	}
	if t == nil {
		fmt.Printf("<nil>\n")
		return
	}
	fmt.Printf("%v", t.val)
	fmt.Println()
	if t.left != nil {
		t.left._print(indent + 2)
	}
	if t.right != nil {
		t.right._print(indent + 2)
	}
}

func (t *TreeNode) Print() {
	t._print(0)
}

func (t *TreeNode) Insert(val int) *TreeNode {
	if t == nil {
		return NewTreeNode(val, nil, nil)
	}
	if val < t.val {
		t.left = t.left.Insert(val)
	} else if val > t.val {
		t.right = t.right.Insert(val)
	}
	return t
}

func (t *TreeNode) Find(val int) *TreeNode {
	if t == nil || t.val == val {
		return t
	}
	if val < t.val {
		return t.left.Find(val)
	} else {
		return t.right.Find(val)
	}
}

func (t *TreeNode) FindFloorAndCeiling(val int, floor *TreeNode, ceil *TreeNode) (*TreeNode, *TreeNode) {
	if t == nil {
		return floor, ceil
	}
	if t.val == val {
		return t, t
	} else if val < t.val {
		return t.left.FindFloorAndCeiling(val, floor, t)
	} else {
		return t.right.FindFloorAndCeiling(val, t, ceil)
	}
}

type BinarySearchTree struct {
	root *TreeNode
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (t *BinarySearchTree) Insert(val int) {
	t.root = t.root.Insert(val)
}

func (t *BinarySearchTree) Find(val int) bool {
	return t.root.Find(val) != nil
}

func (t *BinarySearchTree) Print() {
	t.root.Print()
}

func (t *BinarySearchTree) FindFloorAndCeiling(val int) (*TreeNode, *TreeNode) {
	return t.root.FindFloorAndCeiling(val, nil, nil)
}

func main() {
	t := NewBinarySearchTree()
	t.Insert(7)
	t.Insert(10)
	t.Insert(5)
	t.Insert(-1)
	t.Insert(25)
	t.Insert(6)

	t.Print()

	fmt.Printf("Find val 10: %t\n", t.Find(10))
	fmt.Printf("Find val 11: %t\n", t.Find(11))

	floor, ceil := t.FindFloorAndCeiling(5)
	fmt.Printf("value %v, floor: %v, ceiling: %v\n", 5, floor, ceil)

	floor, ceil = t.FindFloorAndCeiling(7)
	fmt.Printf("value %v, floor: %v, ceiling: %v\n", 7, floor, ceil)

	floor, ceil = t.FindFloorAndCeiling(8)
	fmt.Printf("value %v, floor: %v, ceiling: %v\n", 8, floor, ceil)

	floor, ceil = t.FindFloorAndCeiling(4)
	fmt.Printf("value %v, floor: %v, ceiling: %v\n", 4, floor, ceil)

	floor, ceil = t.FindFloorAndCeiling(-4)
	fmt.Printf("value %v, floor: %v, ceiling: %v\n", -4, floor, ceil)
}
