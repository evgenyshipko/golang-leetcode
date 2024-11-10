package palindrome

// https://leetcode.com/problems/palindrome-number/description/

import "fmt"

func IsPalindrome(x int) bool {
	strX := fmt.Sprint(x)
	length := len(strX)
	runes := []rune(strX)
	for i := 0; i < length; i++ {
		if runes[i] != runes[length-i-1] {
			return false
		}
	}
	return true
}
