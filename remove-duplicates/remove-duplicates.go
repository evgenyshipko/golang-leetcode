package remove_duplicates

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

import "fmt"

func RemoveDuplicates(nums []int) int {

	current := nums[0]
	storageSlice := make([]int, 0, len(nums))
	storageSlice = append(storageSlice, current)

	for _, n := range nums {
		if current != n {
			current = n
			storageSlice = append(storageSlice, n)
		}
	}

	for i := 0; i < len(storageSlice); i++ {
		nums[i] = storageSlice[i]
	}

	fmt.Println(nums)

	return len(storageSlice)
}
