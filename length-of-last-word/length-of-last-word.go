package length_of_last_word

func lengthOfLastWord(s string) int {

	if len(s) == 0 {
		return 0
	}

	for s[len(s)-1] == 32 {
		if len(s)-1 == 0 {
			s = ""
			break
		}
		s = s[:len(s)-1]
	}

	if len(s) == 1 {
		return 1
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 32 {
			return len(s) - 1 - i
		}
	}
	return len(s)
}
