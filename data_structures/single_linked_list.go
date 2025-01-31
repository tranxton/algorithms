package data_structures

type SingleLinkedList struct {
	head *ListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type DoubleLinkedList struct {
	head *BidirectionalNode
	tail *BidirectionalNode
}

type BidirectionalNode struct {
	val  int
	next *BidirectionalNode
	prev *BidirectionalNode
}

func NewSingleLinkedList() SingleLinkedList {
	return SingleLinkedList{}
}

func (list *SingleLinkedList) AppendToSingleLinkedList(val int) *SingleLinkedList {
	if list == nil {
		list = &SingleLinkedList{
			head: &ListNode{Val: val},
		}

		return list
	}

	list.head = &ListNode{Val: val}

	return list
}

func (list *SingleLinkedList) InsertToSingleLinkedList(val int) *SingleLinkedList {
	if list == nil || list.head == nil {
		return list
	}

	current := list.head

	for current != nil {
		current = current.Next
	}

	current.Next = &ListNode{Val: val}

	return list
}

func (list *SingleLinkedList) InsertToSingleLinkedListAfter(after int, val int) *SingleLinkedList {
	if list == nil || list.head == nil {
		return list
	}

	current := list.head

	for current != nil {
		if current.Val == after {
			break
		}

		current = current.Next
	}

	var next *ListNode
	if current.Next != nil {
		next = current.Next
	}

	current.Next = &ListNode{Val: val, Next: next}

	return list
}

func (list *SingleLinkedList) DeleteFromSingleLinkedList(val int) *SingleLinkedList {
	if list == nil || list.head == nil {
		return list
	}

	current := list.head

	for current != nil {
		if current.Val == val {
			break
		}

		current = current.Next
	}
	var next *ListNode
	if current.Next != nil {
		next = current.Next
	}
	current = next

	return list
}
