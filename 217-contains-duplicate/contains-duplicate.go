package _17_contains_duplicate

// https://leetcode.com/problems/contains-duplicate/

func containsDuplicate(nums []int) bool {
	myMap := map[int]bool{}
	for _, itemI := range nums {
		if myMap[itemI] {
			return true
		}
		myMap[itemI] = true
	}
	return false
}
