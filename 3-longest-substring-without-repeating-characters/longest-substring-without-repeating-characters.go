func lengthOfLongestSubstring(s string) int {
    charSet := make(map[rune]bool)
    left := 0
    maxLength := 0
    runes := []rune(s)

    for right := 0; right < len(runes); right++ {

        for charSet[runes[right]] {
            delete(charSet, runes[left])
            left++
        }

        charSet[runes[right]] = true

        currentLength := right - left + 1
        if currentLength > maxLength {
            maxLength = currentLength
        }
    }

    return maxLength
}