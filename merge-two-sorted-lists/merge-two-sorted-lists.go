package merge_two_sorted_lists

// https://leetcode.com/problems/merge-two-sorted-lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{Val: val, Next: nil}
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 != nil {
		return list2
	}

	if list1 != nil && list2 == nil {
		return list1
	}

	var resultHead *ListNode

	currResult := resultHead
	list1curr := list1
	list2curr := list2

	for list1curr != nil && list2curr != nil {
		if list1curr.Val <= list2curr.Val {
			newNode := NewListNode(list1curr.Val)
			if resultHead == nil {
				resultHead = newNode
				currResult = newNode
			} else {
				currResult.Next = newNode
			}
			list1curr = list1curr.Next
		} else {
			newNode := NewListNode(list2curr.Val)
			if resultHead == nil {
				resultHead = newNode
				currResult = newNode
			} else {
				currResult.Next = newNode
			}
			list2curr = list2curr.Next
		}
		if currResult.Next != nil {
			currResult = currResult.Next
		}
	}

	if list1curr == nil && list2curr != nil {
		currResult.Next = list2curr
	}

	if list1curr != nil && list2curr == nil {
		currResult.Next = list1curr
	}

	return resultHead
}
