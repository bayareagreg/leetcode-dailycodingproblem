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
	for range indent {
		fmt.Printf(" ")
	}
	if t == nil {
		fmt.Printf("<nil>\n")
		return
	}
	switch t.val {
	case '+', '-', '*', '/':
		fmt.Printf("%c", t.val)
	default:
		fmt.Printf("%v", t.val)
	}
	fmt.Println()
	if t.left != nil || t.right != nil {
		printHelper(t.left, indent+2)
		printHelper(t.right, indent+2)
	}
}

func (t *TreeNode) Print() {
	printHelper(t, 0)
}

func eval(t *TreeNode) int {
	if t.left == nil && t.right == nil {
		// leaf node
		return t.val
	}
	if t.left != nil && t.right == nil {
		panic("invalid tree")
	}
	if t.left == nil && t.right != nil {
		panic("invalid tree")
	}
	switch t.val {
	case '+':
		return eval(t.left) + eval(t.right)
	case '-':
		return eval(t.left) - eval(t.right)
	case '*':
		return eval(t.left) * eval(t.right)
	case '/':
		return eval(t.left) / eval(t.right)
	default:
		panic("invalid tree")
	}
}

func main() {
	t := NewTreeNode(
		'*',
		NewTreeNode(
			'+',
			NewTreeNode(3, nil, nil),
			NewTreeNode(2, nil, nil)),
		NewTreeNode(
			'+',
			NewTreeNode(4, nil, nil),
			NewTreeNode(5, nil, nil)))
	t.Print()
	fmt.Printf("eval tree: %v\n", eval(t))
}
