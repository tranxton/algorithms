package leet_code

import (
	"reflect"
	"testing"
)

func TestRemoveElementExcept(t *testing.T) {
	removedElementsNumber := RemoveElementExcept([]int{3, 2, 2, 3}, 2)
	expectedRemovedElementsNumber := 2

	if removedElementsNumber != 2 {
		t.Errorf("expected %d, got %d", expectedRemovedElementsNumber, removedElementsNumber)
	}
}

func TestHasPair(t *testing.T) {
	result := HasPair([]int{1, 2, 3, 1}, 5)
	exceptedResult := true

	if !result {
		t.Errorf("expected %t, got %t", exceptedResult, result)
	}
}

func TestMoveZerosToEnd(t *testing.T) {
	result := MoveZerosToEnd([]int{0, 1, 0, 3, 12})
	exceptedResult := []int{1, 3, 12, 0, 0}

	if !reflect.DeepEqual(result, exceptedResult) {
		t.Errorf("expected %v, got %v", exceptedResult, result)
	}
}

func TestTwoSum(t *testing.T) {
	result := TwoSum([]int{2, 7, 11, 15, 19}, 30)
	exceptedResult := []int{2, 4}

	if !reflect.DeepEqual(result, exceptedResult) {
		t.Errorf("expected %v, got %v", exceptedResult, result)
	}
}

func TestIsPalindromeTwoPointers(t *testing.T) {
	result := IsPalindromeTwoPointers("A man, a plan, a canal: Panama")
	exceptedResult := true

	if !result {
		t.Errorf("expected %t, got %t", exceptedResult, result)
	}
}
