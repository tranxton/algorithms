package leet_code

import (
	"algorithms/data_structures"
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	list1 := data_structures.ListNode{Val: 1, Next: &data_structures.ListNode{Val: 2, Next: &data_structures.ListNode{Val: 4}}}
	list2 := data_structures.ListNode{Val: 1, Next: &data_structures.ListNode{Val: 3, Next: &data_structures.ListNode{Val: 4}}}
	mergedList := MergeTwoLists(&list1, &list2)
	expectedMergedList := data_structures.ListNode{Val: 1, Next: &data_structures.ListNode{Val: 1, Next: &data_structures.ListNode{Val: 2, Next: &data_structures.ListNode{Val: 3, Next: &data_structures.ListNode{Val: 4, Next: &data_structures.ListNode{Val: 4}}}}}}

	if !reflect.DeepEqual(mergedList, &expectedMergedList) {
		t.Errorf("expectedMergedList: %v, got: %v", expectedMergedList.ToString(), mergedList.ToString())
	}
}

func TestReverseList(t *testing.T) {
	list := data_structures.ListNode{Val: 1, Next: &data_structures.ListNode{Val: 2, Next: &data_structures.ListNode{Val: 3, Next: &data_structures.ListNode{Val: 4, Next: &data_structures.ListNode{Val: 5}}}}}
	reversedList := ReverseList(&list)
	expectedReversedList := data_structures.ListNode{Val: 5, Next: &data_structures.ListNode{Val: 4, Next: &data_structures.ListNode{Val: 3, Next: &data_structures.ListNode{Val: 2, Next: &data_structures.ListNode{Val: 1}}}}}

	if !reflect.DeepEqual(reversedList, &expectedReversedList) {
		t.Errorf("expectedReversedList: %v, got: %v", expectedReversedList.ToString(), reversedList.ToString())
	}
}
