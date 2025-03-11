package _3_remove_duplicates_from_sorted_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNodeNilPointer(t *testing.T) {
	var head *ListNode
	assert.Equal(t, head, RemoveDuplicates(head))
}

func TestEmptyNode(t *testing.T) {
	head := &ListNode{}

	assert.Equal(t, head, RemoveDuplicates(head))
	assert.Equal(t, []int{0}, ListToSlice(RemoveDuplicates(head)))
}

func TestRemoveDuplicates(t *testing.T) {

	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 99,
						},
					},
				},
			},
		},
	}

	assert.Equal(t, head, RemoveDuplicates(head))
	assert.Equal(t, []int{1, 2, 3, 99}, ListToSlice(RemoveDuplicates(head)))
}

func TestRemoveDuplicates2(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
			},
		},
	}

	assert.Equal(t, head, RemoveDuplicates(head))
	assert.Equal(t, []int{1}, ListToSlice(RemoveDuplicates(head)))
}
