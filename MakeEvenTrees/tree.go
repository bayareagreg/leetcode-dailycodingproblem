package main

import "fmt"

type TreeNode struct {
	val      int
	children []*TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		val:      val,
		children: make([]*TreeNode, 0),
	}
}

func NewTreeNodeWithChildren(val int, children []*TreeNode) *TreeNode {
	return &TreeNode{
		val:      val,
		children: children,
	}
}

func depthFirstPrintHelper(root *TreeNode, indent string) {
	if root == nil {
		fmt.Printf("%s<nil>\n", indent)
	} else {
		fmt.Printf("%s%v\n", indent, root.val)
		for _, c := range root.children {
			depthFirstPrintHelper(c, indent+"  ")
		}
	}
}

func depthFirstPrint(root *TreeNode) {
	depthFirstPrintHelper(root, "")
}

func depthFirstTraversal(root *TreeNode) {
	if root != nil {
		fmt.Printf("%v,", root.val)
		for _, c := range root.children {
			depthFirstTraversal(c)
		}
	}
}

func breadthFirstTraversal(root *TreeNode) {
	deque := NewDeque()
	deque.append(root)
	for !deque.isEmpty() {
		node := deque.popLeft()
		fmt.Printf("%v,", node.val)
		for _, c := range node.children {
			deque.append(c)
		}
	}
}

func main0() {
	root := NewTreeNodeWithChildren(
		1, []*TreeNode{
			NewTreeNode(2),
			NewTreeNodeWithChildren(
				3, []*TreeNode{
					NewTreeNodeWithChildren(
						4, []*TreeNode{
							NewTreeNode(6),
							NewTreeNode(7),
							NewTreeNode(8),
						}),
					NewTreeNode(5)})})

	//depthFirstTraversal(root)
	breadthFirstTraversal(root)
}
