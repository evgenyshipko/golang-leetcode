package longest_common_prefix

// https://leetcode.com/problems/longest-common-prefix/

func LongestCommonPrefix(strs []string) string {
	counter := 0
	for i := 0; i < len(strs[0]); i++ {
		for _, elem := range strs {
			if len(elem) < i+1 || elem[counter] != strs[0][counter] {
				return strs[0][:counter]
			}
		}
		counter++
	}
	return strs[0][:counter]
}
