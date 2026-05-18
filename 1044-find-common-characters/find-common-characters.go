func commonChars(words []string) []string {
    minFreq := map[rune]int{}

	for _, ch := range words[0] {
		minFreq[ch]++
	}

	for i := 1; i < len(words); i++ {

		currFreq := map[rune]int{}

		for _, ch := range words[i] {
			currFreq[ch]++
		}

		for ch, count := range minFreq {

			if currFreq[ch] < count {
				minFreq[ch] = currFreq[ch]
			}
		}
	}

	result := []string{}

	for ch, count := range minFreq {

		for count > 0 {
			result = append(result, string(ch))
			count--
		}
	}

	return result
}