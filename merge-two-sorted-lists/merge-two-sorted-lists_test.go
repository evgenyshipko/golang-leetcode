package merge_two_sorted_lists

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func generateListNode(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	parentNode := NewListNode(arr[0])
	currentNode := parentNode
	for i := 1; i < len(arr); i++ {
		currentNode.Next = NewListNode(arr[i])
		currentNode = currentNode.Next
	}
	return parentNode
}

func generateSlice(list *ListNode) []int {
	result := make([]int, 0)
	tmp := list
	for tmp != nil {
		result = append(result, tmp.Val)
		tmp = tmp.Next
	}
	return result
}

func TestMergeTwoLists(t *testing.T) {

	list1 := generateListNode([]int{1, 6, 8})
	list2 := generateListNode([]int{2, 3, 4, 10, 11})

	mergedList := MergeTwoLists(list1, list2)

	actual := generateSlice(mergedList)
	expected := []int{1, 2, 3, 4, 6, 8, 10, 11}

	assert.Equal(t, true, reflect.DeepEqual(actual, expected))
}

func TestMergeTwoLists1(t *testing.T) {

	list1 := generateListNode([]int{1, 2, 4})
	list2 := generateListNode([]int{1, 3, 4})

	mergedList := MergeTwoLists(list1, list2)

	actual := generateSlice(mergedList)
	expected := []int{1, 1, 2, 3, 4, 4}

	assert.Equal(t, true, reflect.DeepEqual(actual, expected))
}

func TestMergeTwoListsCorner(t *testing.T) {

	list1 := generateListNode([]int{})
	list2 := generateListNode([]int{0})

	mergedList := MergeTwoLists(list1, list2)

	actual := generateSlice(mergedList)
	expected := []int{0}

	assert.Equal(t, true, reflect.DeepEqual(actual, expected))
}
