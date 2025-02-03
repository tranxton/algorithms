package leet_code

import (
	"reflect"
	"testing"
)

func TestRightSideViewDfs(t *testing.T) {
	binaryTree := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}}}, Right: &TreeNode{Val: 3}}
	rightSideView := RightSideViewDfs(binaryTree)
	expectedRightSideView := []int{1, 3, 4, 5}

	if !reflect.DeepEqual(rightSideView, expectedRightSideView) {
		t.Errorf("expected %v, got %v", expectedRightSideView, rightSideView)
	}
}
func TestRightSideViewBfs(t *testing.T) {
	binaryTree := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}}}, Right: &TreeNode{Val: 3}}
	rightSideView := RightSideViewBfs(binaryTree)
	expectedRightSideView := []int{1, 3, 4, 5}

	if !reflect.DeepEqual(rightSideView, expectedRightSideView) {
		t.Errorf("expected %v, got %v", expectedRightSideView, rightSideView)
	}
}

func TestIsValidBST(t *testing.T) {
	bst := &TreeNode{Val: 5, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 7}}}
	isValid := IsValidBST(bst)
	expectedIsValid := false

	if isValid != expectedIsValid {
		t.Errorf("expected %t, got %t", expectedIsValid, isValid)
	}
}
