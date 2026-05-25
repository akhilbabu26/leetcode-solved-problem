import "strings"

func findWords(words []string) []string {
    row1 := "qwertyuiop"
	row2 := "asdfghjkl"
	row3 := "zxcvbnm"

	var result []string

	for _, word := range words {

		lower := strings.ToLower(word)

		if check(lower, row1) ||
			check(lower, row2) ||
			check(lower, row3) {

			result = append(result, word)
		}
	}

	return result
}

func check(word string, row string) bool {

	for _, ch := range word {

		if !strings.ContainsRune(row, ch) {
			return false
		}
	}

	return true
}