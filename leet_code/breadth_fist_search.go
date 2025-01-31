package leet_code

import (
	"container/list"
	"math"
)

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	pQueue, qQueue := list.New(), list.New()
	pQueue.PushBack(p)
	qQueue.PushBack(q)

	for pQueue.Len() > 0 && qQueue.Len() > 0 {
		pNode := pQueue.Remove(pQueue.Front()).(*TreeNode)
		qNode := qQueue.Remove(qQueue.Front()).(*TreeNode)

		if pNode == nil && qNode == nil {
			continue
		}

		if pNode == nil || qNode == nil || pNode.Val != qNode.Val {
			return false
		}

		pQueue.PushBack(pNode.Left)
		pQueue.PushBack(pNode.Right)
		qQueue.PushBack(qNode.Left)
		qQueue.PushBack(qNode.Right)
	}

	return pQueue.Len() == 0 && qQueue.Len() == 0
}

func IsSymmetricBfs(root *TreeNode) bool {
	queue := list.New()

	queue.PushFront(root.Left)
	queue.PushBack(root.Right)

	for queue.Len() > 0 {
		left := queue.Remove(queue.Front()).(*TreeNode)
		right := queue.Remove(queue.Back()).(*TreeNode)

		if left == nil && right == nil {
			continue
		}

		if left == nil || right == nil || left.Val != right.Val {
			return false
		}

		queue.PushFront(left.Right)
		queue.PushFront(left.Left)
		queue.PushBack(right.Left)
		queue.PushBack(right.Right)
	}

	return true
}

func IsSymmetricDfs(root *TreeNode) bool {
	return isSymmetricTreesDfs(root.Left, root.Right)
}

func isSymmetricTreesDfs(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil || left.Val != right.Val {
		return false
	}

	return isSymmetricTreesDfs(left.Left, right.Right) && isSymmetricTreesDfs(left.Right, right.Left)
}

func MaxDepthDfs(root *TreeNode, max int) int {
	if root == nil {
		return max
	}

	max++

	maxLeft := MaxDepthDfs(root.Left, max)
	maxRight := MaxDepthDfs(root.Right, max)

	if maxLeft > maxRight {
		return maxLeft
	}
	return maxRight
}

func MaxDepthBfs(root *TreeNode) int {
	queue := list.New()
	queue.PushBack(root)
	maxTreeDepth := -1

	for ; queue.Len() > 0; maxTreeDepth++ {
		queueLen := queue.Len()
		for i := 0; i < queueLen; i++ {
			current := queue.Remove(queue.Front()).(*TreeNode)
			if current == nil {
				continue
			}

			queue.PushBack(current.Left)
			queue.PushBack(current.Right)
		}
	}

	return maxTreeDepth
}

func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	_, isBalanceTree := maxDepth(root, 0)

	return isBalanceTree
}

func maxDepth(root *TreeNode, max int) (int, bool) {
	if root == nil {
		return max, true
	}

	max++

	leftDepth, isBalancedLeft := maxDepth(root.Left, max)
	rightDepth, isBalancedRight := maxDepth(root.Right, max)
	isBalancedTrees := isBalancedLeft && isBalancedRight && math.Abs(float64(leftDepth-rightDepth)) <= 1

	if leftDepth > rightDepth {
		return leftDepth, isBalancedTrees
	}

	return rightDepth, isBalancedTrees
}

func MinDepthDfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := MinDepthDfs(root.Left)
	rightDepth := MinDepthDfs(root.Right)

	if rightDepth == 0 {
		return leftDepth + 1
	}

	if leftDepth == 0 {
		return rightDepth + 1
	}

	if leftDepth < rightDepth {
		return leftDepth + 1
	}

	return rightDepth + 1
}

func MinDepthBfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := list.New()
	min := 0
	queue.PushBack(root)

	for queue.Len() > 0 {
		queueLen := queue.Len()
		for i := 0; i < queueLen; i++ {
			current := queue.Remove(queue.Front()).(*TreeNode)
			if current == nil {
				continue
			}

			if current.Left == nil && current.Right == nil {
				return min + 1
			}

			queue.PushBack(current.Left)
			queue.PushBack(current.Right)
		}
		min++
	}

	return min
}

func HasPathSumDfs(root *TreeNode, targetSum int) bool {
	return hasPathSumDfs(root, 0, targetSum)
}

func hasPathSumDfs(root *TreeNode, currentSum int, targetSum int) bool {
	if root == nil {
		return false
	}

	currentSum += root.Val

	if currentSum == targetSum && root.Left == nil && root.Right == nil {
		return true
	}

	return hasPathSumDfs(root.Left, currentSum, targetSum) || hasPathSumDfs(root.Right, currentSum, targetSum)
}

func HasPathSumBfs(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	nodesQueue, sumsQueue := list.New(), list.New()
	nodesQueue.PushBack(root)
	sumsQueue.PushBack(root.Val)

	for nodesQueue.Len() > 0 {
		currentNode := nodesQueue.Remove(nodesQueue.Front()).(*TreeNode)
		currentSum := sumsQueue.Remove(sumsQueue.Front()).(int)
		if currentNode == nil {
			continue
		}

		if currentSum == targetSum && currentNode.Left == nil && currentNode.Right == nil {
			return true
		}

		if currentNode.Left != nil {
			nodesQueue.PushBack(currentNode.Left)
			sumsQueue.PushBack(currentSum + currentNode.Left.Val)
		}

		if currentNode.Right != nil {
			nodesQueue.PushBack(currentNode.Right)
			sumsQueue.PushBack(currentSum + currentNode.Right.Val)
		}
	}

	return false
}

func InvertTreeDfs(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = InvertTreeDfs(root.Right), InvertTreeDfs(root.Left)

	return root
}

func InvertTreeBfs(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	treeQueue := list.New()
	treeQueue.PushFront(root)

	for treeQueue.Len() > 0 {
		currentTreeNode := treeQueue.Remove(treeQueue.Front()).(*TreeNode)
		if currentTreeNode == nil {
			continue
		}

		currentTreeNode.Left, currentTreeNode.Right = currentTreeNode.Right, currentTreeNode.Left

		treeQueue.PushFront(currentTreeNode.Right)
		treeQueue.PushFront(currentTreeNode.Left)
	}

	return root
}

func SumOfLeftLeavesDfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		sum += root.Left.Val
	}

	return sum + SumOfLeftLeavesDfs(root.Left) + SumOfLeftLeavesDfs(root.Right)
}

func SumOfLeftLeavesBfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := list.New()
	queue.PushBack(root)
	sum := 0

	for queue.Len() > 0 {
		currentNode := queue.Remove(queue.Front()).(*TreeNode)
		if currentNode == nil {
			continue
		}

		if currentNode.Left != nil && currentNode.Left.Left == nil && currentNode.Left.Right == nil {
			sum += currentNode.Left.Val
		}

		queue.PushBack(currentNode.Left)
		queue.PushBack(currentNode.Right)
	}

	return sum
}
