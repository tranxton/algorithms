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

func TestMaxDistanceToNearestPerson(t *testing.T) {
	//seats := []int{1, 0, 0, 0, 1}
	//expectedMaxDistanceToNearestPerson := 2
	//seats := []int{1, 0, 1, 0, 0, 1, 0, 0, 0, 1}
	//expectedMaxDistanceToNearestPerson := 2
	seats := []int{1, 0, 1, 0, 0, 0}
	expectedMaxDistanceToNearestPerson := 3
	maxDistanceToNearestPerson := MaxDistanceToNearestPerson(seats)

	if maxDistanceToNearestPerson != expectedMaxDistanceToNearestPerson {
		t.Errorf("expected %d, got %d", expectedMaxDistanceToNearestPerson, maxDistanceToNearestPerson)
	}
}

func TestLongestMonotonicSubarray(t *testing.T) {
	//arr := []int{2, 7, 5, 4, 4, 3}
	//expectedLongestMonotonicSubarray := []int{1, 3}
	//arr := []int{1, 1}
	//expectedLongestMonotonicSubarray := []int{0, 0}
	//arr := []int{1, 2, 3, 4, 5}
	//expectedLongestMonotonicSubarray := []int{0, 4}
	//arr := []int{5, 4, 3, 2, 1}
	//expectedLongestMonotonicSubarray := []int{0, 4}
	//arr := []int{1, 3, 5, 4, 2, 0}
	//expectedLongestMonotonicSubarray := []int{2, 5}
	//arr := []int{}
	//expectedLongestMonotonicSubarray := []int{0, 0}
	arr := []int{1, 3, 5, 4, 2, 2, 2, 2, 2, 2, 0}
	expectedLongestMonotonicSubarray := []int{0, 2}
	longestMonotonicSubarray := LongestMonotonicSubarray(arr)

	if !reflect.DeepEqual(longestMonotonicSubarray, expectedLongestMonotonicSubarray) {
		t.Errorf("expected %v, got %v", expectedLongestMonotonicSubarray, longestMonotonicSubarray)
	}
}

func TestMakeRanges(t *testing.T) {
	//arr := []int{}
	//expectedResult := ""
	//arr := []int{1}
	//expectedResult := "1"
	//arr := []int{5, 1, 3, 4, 7, 9, 8, 14}
	//expectedResult := "1,3->5,7->9,14"
	//arr := []int{5, 1, 3, 7}
	//expectedResult := "1,3,5,7"
	arr := []int{7, 5, 1, 3, 4, 2, 6, 8, 11, 10, 9}
	expectedResult := "1->11"
	result := MakeRanges(arr)

	if result != expectedResult {
		t.Errorf("expected %s, got %s", expectedResult, result)
	}
}
