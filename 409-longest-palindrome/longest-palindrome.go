func longestPalindrome(s string) int {
    count := make(map[rune]int)

	for _, ch := range s {
		count[ch]++
	}

	length := 0
	hasOdd := false

	for _, freq := range count {

		if freq%2 == 0 {
			length += freq
		} else {

			length += freq - 1
			hasOdd = true
		}
	}

	if hasOdd {
		length++
	}

	return length
}