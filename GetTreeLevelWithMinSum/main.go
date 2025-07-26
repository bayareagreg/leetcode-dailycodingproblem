package main

import (
	"fmt"
	"math"
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

type Queue struct {
	queue [][]int
}

func NewQueue() *Queue {
	return &Queue{
		make([][]int, 0),
	}
}

func (q *Queue) numLevels() int {
	return len(q.queue)
}

func (q *Queue) addLevel() {
	newlevel := make([]int, 0)
	q.queue = append(q.queue, newlevel)
}

func (q *Queue) addValue(level int, val int) {
	q.queue[level] = append(q.queue[level], val)
}

func (q *Queue) print() {
	fmt.Printf("queue: %v\n", q.queue)
}

func (q *Queue) getMinSumLevel() int {
	minSumIndex := -1
	minSum := math.MaxInt32
	for i := range q.queue {
		curLevelSum := 0
		for j := range q.queue[i] {
			curLevelSum += q.queue[i][j]
		}
		if curLevelSum < minSum {
			minSumIndex = i
			minSum = curLevelSum
		}
	}
	return minSumIndex
}

// same as breadth first
func levelOrderTraversal(t *TreeNode, level int, q *Queue) {
	if t == nil {
		return
	}
	if q.numLevels() <= level {
		q.addLevel()
	}
	q.addValue(level, t.val)
	levelOrderTraversal(t.left, level+1, q)
	levelOrderTraversal(t.right, level+1, q)
}

func treeLevelWithMinSum(t *TreeNode) int {
	queue := NewQueue()
	levelOrderTraversal(t, 0, queue)
	queue.print()
	fmt.Printf("tree level with min sum: %v\n", queue.getMinSumLevel())
	return 0
}

func main() {
	t := NewTreeNode(
		1,
		NewTreeNode(2, nil, nil),
		NewTreeNode(
			3,
			NewTreeNode(4, nil, nil),
			NewTreeNode(5, nil, nil)))
	t.Print()
	treeLevelWithMinSum(t)
}
