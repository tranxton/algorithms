package leet_code

import (
	"algorithms/data_structures"
)

func MergeTwoLists(list1 *data_structures.ListNode, list2 *data_structures.ListNode) *data_structures.ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val > list2.Val {
		list2.Next = MergeTwoLists(list1, list2.Next)

		return list2
	}

	list1.Next = MergeTwoLists(list2, list1.Next)

	return list1
}

func ReverseList(head *data_structures.ListNode) *data_structures.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := ReverseList(head.Next)

	head.Next.Next = head
	head.Next = nil

	return newHead
}
