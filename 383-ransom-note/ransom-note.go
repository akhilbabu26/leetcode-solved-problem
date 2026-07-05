func canConstruct(ransomNote string, magazine string) bool {
    count := make([]int, 26)

    for _, ch := range magazine {
        count[ch-'a']++
    }

    for _, ch := range ransomNote {
        if count[ch-'a'] == 0 {
            return false
        }
        count[ch-'a']--
    }

    return true
}