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
	if t.left != nil || t.right != nil {
		t.left._print(indent + 2)
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

func (t *TreeNode) Min() *TreeNode {
	if t == nil {
		return nil
	}
	for t.left != nil {
		t = t.left
	}
	return t
}

func (t *TreeNode) Delete(val int) *TreeNode {
	if t == nil {
		return nil
	}

	if val < t.val {
		t.left = t.left.Delete(val)
	} else if val > t.val {
		t.right = t.right.Delete(val)
	} else {
		if t.left == nil {
			return t.right
		} else if t.right == nil {
			return t.left
		}
		t.val = t.right.val
		t.right = t.right.right
		//minRight := t.right.Min()
		//t.val = minRight.val
		//t.right = t.right.Delete(minRight.val)
	}
	return t
}

type BinarySearchTree struct {
	root *TreeNode
}

func NewBinarySearchTree(root *TreeNode) *BinarySearchTree {
	return &BinarySearchTree{
		root: root,
	}
}

func (t *BinarySearchTree) Insert(val int) {
	t.root = t.root.Insert(val)
}

func (t *BinarySearchTree) Delete(val int) {
	t.root = t.root.Delete(val)
}

func (t *BinarySearchTree) Find(val int) bool {
	return t.root.Find(val) != nil
}

func (t *BinarySearchTree) Print() {
	t.root.Print()
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

func constructAllBSTs(n int) []*BinarySearchTree {
	nodes := make_trees(1, n)
	trees := make([]*BinarySearchTree, 0)
	for i := range nodes {
		trees = append(trees, NewBinarySearchTree(nodes[i]))
	}
	return trees
}

func main() {
	tt := constructAllBSTs(5)
	for i := range tt {
		tt[i].Print()
		fmt.Println("-----------------")
	}
	fmt.Printf("Total trees: %v\n", len(tt))
}
