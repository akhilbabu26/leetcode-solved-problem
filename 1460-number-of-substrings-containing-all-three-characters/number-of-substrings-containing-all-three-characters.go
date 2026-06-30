func numberOfSubstrings(s string) int {
    n := len(s)
    freq := make([]int, 3)

    left := 0
    ans := 0

    for right := 0; right < n; right++ {
        freq[s[right]-'a']++

        for freq[0] > 0 && freq[1] > 0 && freq[2] > 0 {
            ans += n - right
            freq[s[left]-'a']--
            left++
        }
    }

    return ans
}