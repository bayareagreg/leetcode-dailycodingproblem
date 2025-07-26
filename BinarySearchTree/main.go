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

func main() {
	t := NewBinarySearchTree(nil)
	t.Insert(7)
	t.Insert(10)
	t.Insert(5)
	t.Insert(-1)
	t.Insert(25)
	t.Insert(6)

	t.Print()

	fmt.Printf("Find val 10: %t\n", t.Find(10))
	fmt.Printf("Find val 11: %t\n", t.Find(11))

	t.Delete(-1)
	t.Print()
}
