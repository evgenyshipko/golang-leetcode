package _3_remove_duplicates_from_sorted_list

/*
Input: head = [1,1,2,3,3]
Output: [1,2,3]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveDuplicates(head *ListNode) *ListNode {
	currentList := head
	for {
		if currentList == nil || currentList.Next == nil {
			break
		}

		if currentList.Next.Val == currentList.Val {
			currentList.Next = currentList.Next.Next
			continue
		}
		currentList = currentList.Next
	}
	return head
}

func ListToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
