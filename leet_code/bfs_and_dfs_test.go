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

func TestFindSimilarRootsDfs(t *testing.T) {
	/*root := &TreeNodeRune{'A',
		&TreeNodeRune{'C',
			&TreeNodeRune{'A', &TreeNodeRune{'B', nil, nil}, nil},
			&TreeNodeRune{'D', nil, nil},
		},
		&TreeNodeRune{'B',
			&TreeNodeRune{'A', nil, nil},
			&TreeNodeRune{'D', nil, &TreeNodeRune{'C', nil, nil}},
		},
	}
	expectedKey, expectedSets := "ABCD", []string{"CABD", "BACD", "AABCDABCD"}*/
	/*root := &TreeNodeRune{
		Val:   'A',
		Left:  &TreeNodeRune{Val: 'C'},
		Right: &TreeNodeRune{Val: 'C'},
	}
	expectedKey, expectedSets := "C", []string{"C", "C"}*/
	/*root := &TreeNodeRune{
		Val: 'A',
	}
	expectedKey, expectedSets := "", []string{}*/
	/*root := &TreeNodeRune{}
	expectedKey, expectedSets := "", []string{}*/
	root := &TreeNodeRune{'A',
		&TreeNodeRune{'A',
			&TreeNodeRune{'A', &TreeNodeRune{'A', nil, nil}, nil},
			&TreeNodeRune{'A', nil, nil},
		},
		&TreeNodeRune{'A',
			&TreeNodeRune{'A', nil, nil},
			&TreeNodeRune{'A', nil, &TreeNodeRune{'A', nil, nil}},
		},
	}
	expectedKey, expectedSets := "A", []string{"A", "AA", "A", "AAA", "A", "A", "AA", "AAA", "AAA"}

	result := FindSimilarRootsDfs(root)

	for key, sets := range result {
		if key != expectedKey {
			t.Errorf("expected key %s, got %s", expectedSets, sets)
		}

		if !reflect.DeepEqual(sets, expectedSets) {
			t.Errorf("expected sets %s, got %s", expectedSets, sets)
		}
	}
}
