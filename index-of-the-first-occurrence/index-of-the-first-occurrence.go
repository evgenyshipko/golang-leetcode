package index_of_the_first_occurrence

func IndexFirstOccur(haystack string, needle string) int {
	for index, letter := range haystack {
		if letter == rune(needle[0]) &&
			index+len(needle) <= len(haystack) &&
			haystack[index:index+len(needle)] == needle {
			return index
		}
	}
	return -1
}
